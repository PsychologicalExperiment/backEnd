build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
	scp experiment_server Javis:/data/server/
	ssh Javis "cd /data/server/ && ./experiment_server"