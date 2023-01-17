
```shell
protoc -I=./ \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  --govalidators_out=paths=source_relative:. \
  --go-grpc_opt=require_unimplemented_servers=false \
  experiment_server/experiment_server.proto
```