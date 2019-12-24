#!/bin/bash

script_path=$(cd `dirname $0`; pwd)
echo '脚本执行目录:'$script_path;
echo '编译linux/amd64可执行文件'
$script_path/gox -osarch "linux/amd64" -output $script_path/../build/kratos-demo  $script_path/../cmd

