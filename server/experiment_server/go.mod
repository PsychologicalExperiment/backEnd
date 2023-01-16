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
	github.com/gofrs/uuid v4.3.1+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/prometheus/client_golang v1.14.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	google.golang.org/grpc v1.51.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gorm.io/driver/mysql v1.4.5
	gorm.io/gorm v1.24.3
)
