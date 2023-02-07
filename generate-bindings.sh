#!/bin/sh
# Copyright 2023 Stichting ThingsIX Foundation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

set -e

generate_bindings() {
  local pkg=$1
  local abi_file="IMapperRegistry.abi"  

  echo Generate bindings for package $pkg
  abigen --abi "$abi_file" --pkg $pkg --out "bindings.go"

  echo Delete work directory
  rm -rf $work_dir
}

# generate go-bindings for router registry contract
generate_bindings mapper_registry

# install dependencies
echo Cleanup go dependencies
go mod tidy
