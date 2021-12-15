RUN_NAME="client"
mkdir -p output/bin
export GO111MODULE=on
go build  -o output/bin/${RUN_NAME}