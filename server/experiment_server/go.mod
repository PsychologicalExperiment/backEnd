module github.com/PsychologicalExperiment/backEnd/server/experiment_server

go 1.16

replace (
	github.com/PsychologicalExperiment/backEnd/api/experiment_server => ../../api/experiment_server
	github.com/PsychologicalExperiment/backEnd/api/user_info_server => ../../api/user_info_server
	github.com/PsychologicalExperiment/backEnd/util => ../../util
	github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
)

require (
	github.com/PsychologicalExperiment/backEnd v0.0.3
	github.com/PsychologicalExperiment/backEnd/api/experiment_server v0.0.0-00010101000000-000000000000
	github.com/PsychologicalExperiment/backEnd/api/user_info_server v0.0.0-00010101000000-000000000000
	github.com/PsychologicalExperiment/backEnd/util v0.0.0-00010101000000-000000000000
	github.com/gofrs/uuid v4.3.1+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/mwitkow/go-proto-validators v0.3.2 // indirect
	go.etcd.io/etcd/client/v3 v3.5.6
	google.golang.org/grpc v1.52.0
	gorm.io/driver/mysql v1.4.5
	gorm.io/gorm v1.24.3
)
