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
#RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/g/golang-1.6/golang-1.6-go_1.6-0ubuntu1~14.04_amd64.deb
RUN wget https://storage.googleapis.com/k8s-bandwidth-test/golang-1.6-go_1.6-0ubuntu1~14.04_amd64.deb
#RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/g/golang-1.6/golang-1.6-src_1.6-0ubuntu1~14.04_amd64.deb
RUN wget https://storage.googleapis.com/k8s-bandwidth-test/golang-1.6-src_1.6-0ubuntu1~14.04_amd64.deb
RUN dpkg -i golang-1.6-*.deb

RUN echo "PATH=$PATH:/usr/lib/go-1.6/bin" >> /env.sh

ENTRYPOINT [ "/bin/bash" ]
