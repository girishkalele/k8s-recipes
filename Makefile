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

# Top-level makefile to recursively build everything in this repository
all: build

build:
	for i in `find . -mindepth 2 -name Makefile`; do pushd `dirname $$i`; make build; popd; done

run: build
	docker run $(IMAGE)

push: build
	for i in `find . -mindepth 2 -name Makefile`; do pushd `dirname $$i`; make push; popd; done

