/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// A UDP client to test bounce-backs from pods behind a LoadBalancer service.

// Methodology:
//  Deploy the UDP server which will respond to UDP pings with our pod name that was exposed via
//  the Downward API as an environment variable.
//  The client will aggregate results and measure the distribution over all the backend pods behind the
//  LB service. Depending on the number of replicas and number of nodes, we should be able to measure
//  unbalances with the ESIPP changes.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"
)

var externalIp string
var lock sync.Mutex

// Print the distribution every N seconds
func printDistribution(dmap map[string]int) {
	for {
		time.Sleep(time.Second * 5)
		var total float64
		lock.Lock()
		keys := []string{}
		for k, n := range dmap {
			total = total + float64(n)
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := dmap[k]
			var nf float64 = float64(n)
			fmt.Printf("%s : %3.2f%% (%d/%d)\n", k, float64(nf*100.0)/total, n, int(total))
			delete(dmap, k) // clear on read
		}
		lock.Unlock()
		fmt.Printf("------------------------\n")
	}
}

func main() {

	flag.StringVar(&externalIp, "ip", "", "LB external IP")
	flag.Parse()

	if len(externalIp) == 0 {
		log.Fatalf("Cannot have blank external IP address")
	}

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:10001", externalIp))
	if err != nil {
		log.Fatalf("Failed to lookup server hostname : logger %s", err)
	}
	laddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:0")
	i := 0

	dmap := make(map[string]int)

	go printDistribution(dmap)

	max := 16
	channels := make([]chan bool, max)
	for i := 0; i < max; i++ {
		channels[i] = make(chan bool)
		go func(channel chan bool) {
			recvBuffer := make([]byte, 8192)
			for {
				conn, err := net.DialUDP("udp", laddr, addr)
				if err != nil {
					log.Fatalf("Failed to setup udp connection (%s)", err)
				}
				msg := strconv.Itoa(i)
				i++
				buf := []byte(msg)
				//fmt.Printf("Sending udp packet %d\n", i)
				_, err = conn.Write(buf)
				if err != nil {
					fmt.Println(msg, err)
				}
				time.Sleep(time.Millisecond * 10)
				//fmt.Printf("Waiting for a response")
				conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200)) // us-east should be around 40 ms away
				n, _, err := conn.ReadFromUDP(recvBuffer)
				if n > 0 && err == nil {
					podname := string(recvBuffer[0 : n-1])
					lock.Lock()
					count := dmap[podname]
					dmap[podname] = count + 1
					lock.Unlock()
					//fmt.Printf("%s responded\n", podname)
				} else if err != nil {
					if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
						// time outs are ok, add to the "timeout" phantom pod
						lock.Lock()
						dmap["timeout"] = 1 + dmap["timeout"]
						lock.Unlock()
					}
				}
				conn.Close()
			}
		}(channels[i])
	}
	for i = 0; i < max; i++ {
		<-channels[i]
	}
}
