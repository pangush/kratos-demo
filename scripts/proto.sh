#!/usr/bin/env bash

script_path=$(cd `dirname $0`; pwd)
echo '脚本目录:'$script_path;
protoc --go_out=plugins=grpc:$script_path/../api -I=$script_path/../api -I=$script_path/../third_party/ $script_path/../api/*.proto