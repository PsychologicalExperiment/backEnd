APP=experiment_server
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP} main.go
	ssh Javis-root "sudo rm -rf /data/server/${APP} &> /dev/null"
	ssh -q Javis-root "kill -9 `ps -ef | grep -v grep | grep ${APP} | awk '{print $$2}'` > /dev/null 2>&1 &"
	scp bin/${APP} Javis-root:/data/server/
	ssh Javis-root "cd /data/server/ && chmod 777 ${APP} && ./${APP} > /dev/null 2>&1 &"