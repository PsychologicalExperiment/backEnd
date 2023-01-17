APP=experiment_server
build:
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP} main.go

javis:
	ssh javis "sudo rm -rf /data/psychology/${APP}/${APP} &> /dev/null"
	scp bin/${APP} javis:/data/psychology/${APP}/${APP}
	scp -r config javis:/data/psychology/${APP}/config
	scp -r restart.sh javis:/data/psychology/${APP}/
	ssh javis "cd /data/psychology/${APP} && sh restart.sh ${APP} &> /dev/null 2>&1 &"

ahan:
	ssh ahan "sudo rm -rf /data/psychology/${APP}/${APP} &> /dev/null"
	scp bin/${APP} ahan:/data/psychology/${APP}/${APP}
	scp -r config ahan:/data/psychology/${APP}/config
	scp -r restart.sh ahan:/data/psychology/${APP}/
	ssh ahan "cd /data/psychology/${APP} && sh restart.sh &> /dev/null 2>&1 &"

musk:
	ssh musk "sudo rm -rf /data/psychology/${APP}/${APP} &> /dev/null"
	scp bin/${APP} musk:/data/psychology/${APP}/${APP}
	scp -r config ahan:/data/psychology/${APP}/config
	scp -r restart.sh musk:/data/psychology/${APP}/
	ssh musk "cd /data/psychology/${APP} && sh restart.sh &> /dev/null 2>&1 &"

## 更新idl
idl:
	cd ../../api/ && protoc -I=./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false experiment_server/experiment_server.proto

.PHONY:
	build deploy idl