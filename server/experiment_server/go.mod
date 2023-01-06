module github.com/PsychologicalExperiment/backEnd/server/experiment_server

go 1.16

replace (
	github.com/PsychologicalExperiment/backEnd/util => ../../util
	github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
)

require (
	github.com/PsychologicalExperiment/backEnd v0.0.3
	github.com/PsychologicalExperiment/backEnd/api/experiment_server v0.0.0-20230106045604-ee996205ec1e
	github.com/PsychologicalExperiment/backEnd/util v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/prometheus/client_golang v1.14.0
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.51.0
	gopkg.in/validator.v2 v2.0.1
	gorm.io/driver/mysql v1.4.5
	gorm.io/gorm v1.24.3
)
