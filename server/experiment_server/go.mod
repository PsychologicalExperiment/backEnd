module github.com/PsychologicalExperiment/backEnd/server/experiment_server

go 1.16

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/PsychologicalExperiment/backEnd v0.0.0-20221225123159-ddef9bc70e4a
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/prometheus/client_golang v1.14.0
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.51.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/validator.v2 v2.0.1
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)
