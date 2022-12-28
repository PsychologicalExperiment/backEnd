APP=experiment_server
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP} main.go
	ssh Javis-root "sudo rm -rf /data/server/${APP}"
	scp experiment_server Javis-root:/data/server/
	ssh Javis-root "cd /data/server/ && chmod 777 ${APP} && ./${APP}"