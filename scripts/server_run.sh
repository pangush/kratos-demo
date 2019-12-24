script_path=$(cd `dirname $0`; pwd)

echo '脚本执行目录:'$script_path;
go run $script_path/../cmd/main.go -conf $script_path/../configs -http.perf tcp://0.0.0.0:2306