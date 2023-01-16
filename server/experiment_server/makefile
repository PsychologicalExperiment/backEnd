APP=experiment_server
build:
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP} main.go

deploy:
	ssh ${HOST} "sudo rm -rf /data/psychology/${APP} &> /dev/null"
	scp bin/${APP} ${HOST}:/data/psychology/${APP}
	ssh ${HOST} "cd /data/psychology/${APP} && chmod 777 ${APP} && sh restart.sh &> /dev/null 2>&1 &"

## 更新idl
idl:
	protoc -I=../../api/ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. \api/experiment_server/experiment_server.proto

.PHONY:
	build deploy idl