#!/bin/bash
#
# Copyright 2022 Authors of Cilium
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

# Simple script to make sure viper.GetStringMapString should not be used.
# Related upstream issue https://github.com/spf13/viper/issues/911
if grep -r --exclude-dir={.git,_build,vendor,contrib} -i --include \*.go "viper.GetStringMapString" .; then
  echo "Found viper.GetStringMapString(key) usage. Please use command.GetStringMapString(viper.GetViper(), key) instead";
  exit 1
fi

if grep -r --exclude-dir={.git,_build,vendor,contrib} -i --include \*.go "StringToStringVar" .; then
  echo "Found flags.StringToStringVar usage. Please use option.NewNamedMapOptions instead";
  exit 1
fi