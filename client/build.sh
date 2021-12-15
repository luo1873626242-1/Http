RUN_NAME="client"
mkdir -p output/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.io,direct
go build  -o output/bin/${RUN_NAME}