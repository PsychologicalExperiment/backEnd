```shell
protoc -I=./ \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
    api/experiment_server/experiment_server.proto
```