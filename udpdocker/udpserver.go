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
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		log.Fatalf("Failed to create udp address %s", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on udp port (%s)", err)
	}
	defer conn.Close()

	buf := make([]byte, 8192)
	hostname, _ := os.Hostname()
	hostnameBuf := make([]byte, len(hostname)+1)
	copy(hostnameBuf[:], hostname)
	fmt.Println("Waiting for UDP packets")
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		// Respond back with our hostname
		conn.WriteToUDP(hostnameBuf, addr)
	}
}
