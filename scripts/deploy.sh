#!/bin/bash
script_path=$(cd `dirname $0`; pwd)

bash $script_path/build.sh
ssh root@localhost << eoff
ps aux | grep kratos-demo | awk '{print \$2}' | xargs kill

supervisorctl stop kratos-demo

eoff
scp $script_path/../build/kratos-demo root@39.97.181.40:/web/go/kratos-demo
ssh root@localhost > /dev/null 2>&1 << eoff
echo 1

supervisorctl restart kratos-demo
exit
eoff
