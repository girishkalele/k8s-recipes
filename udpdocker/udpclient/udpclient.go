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

package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
    "log"
)

func main() {
    addr,err := net.ResolveUDPAddr("udp","logger:10001")
    if err != nil {
        log.Fatalf("Failed to lookup server hostname : logger %s", err)
    }
    laddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:0")
    conn, err := net.DialUDP("udp", laddr, addr)
    if err != nil {
        log.Fatalf("Failed to setup udp connection (%s)", err)
    }
    defer conn.Close()
    i := 0
    for {
        msg := strconv.Itoa(i)
        i++
        buf := []byte(msg)
	fmt.Printf("Sending udp packet %d\n", i)
        _,err := conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
        time.Sleep(time.Millisecond * 100)
    }
}
