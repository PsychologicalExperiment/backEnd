module github.com/PsychologicalExperiment/backEnd/server/user_info_server

go 1.19

replace (
	github.com/PsychologicalExperiment/backEnd/api/user_info_server => ../../api/user_info_server
	github.com/PsychologicalExperiment/backEnd/util => ../../util
)

require (
	github.com/PsychologicalExperiment/backEnd v0.0.0-20221228092006-c4424c0032c4
	github.com/PsychologicalExperiment/backEnd/api/user_info_server v0.0.0-20221228092006-c4424c0032c4
	github.com/PsychologicalExperiment/backEnd/util v0.0.0-20221228092006-c4424c0032c4
	github.com/go-redis/redis/v8 v8.11.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	google.golang.org/grpc v1.51.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.6 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.6 // indirect
	go.etcd.io/etcd/client/v3 v3.5.6 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
