APP=experiment_server
build:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP} main.go

deploy:
	ssh Javis-root "sudo rm -rf /data/server/${APP} &> /dev/null"
	scp bin/${APP} Javis-root:/data/server/
	ssh Javis-root "cd /data/server/ && chmod 777 ${APP} && sh restart.sh &> /dev/null 2>&1 &"

.PHONY:
	build deploy