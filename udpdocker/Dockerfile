# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from debian:jessie

RUN apt-get update
RUN apt-get install -y file wget curl tcpdump vim net-tools bridge-utils iptables ebtables dnsutils nmap
RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/g/golang-1.6/golang-1.6-go_1.6-0ubuntu1~14.04_amd64.deb
RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/g/golang-1.6/golang-1.6-src_1.6-0ubuntu1~14.04_amd64.deb
RUN dpkg -i golang-1.6-*.deb
RUN rm -f *.deb
RUN rm -rf /go/src/udpserver
RUN mkdir -p /go/src/udpserver
COPY udpserver.go /go/src/udpserver

RUN find / -name go -type f
RUN PATH=$PATH:/usr/lib/go-1.6/bin go version
RUN GOPATH=/go CGO_ENABLED=0 PATH=$PATH:/usr/lib/go-1.6/bin go build -o /go/src/udpserver/udpserver /go/src/udpserver/udpserver.go 
RUN /usr/bin/file /go/src/udpserver/udpserver
RUN ls -lh /go/src/udpserver/udpserver

ENTRYPOINT [ "/go/src/udpserver/udpserver" ]
