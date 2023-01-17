clean:
	rm -rf experiment_server/*.go user_info_server/*.go

experiment_server:
	protoc -I=./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --govalidators_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false experiment_server/experiment_server.proto

user_info_server:
	protoc -I=./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --govalidators_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false user_info_server/user_info_server.proto
	
.PHONY:
	experiment_server user_info_server clean