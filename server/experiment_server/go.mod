module github.com/PsychologicalExperiment/backEnd/server/experiment_server

replace github.com/PsychologicalExperiment/backEnd/api/experiment_server => ../../api/experiment_server

go 1.16

require (
	github.com/PsychologicalExperiment/backEnd v0.0.0-20220405085841-41916c8ea5cb // indirect
	github.com/PsychologicalExperiment/backEnd/api/experiment_server v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.0
	google.golang.org/grpc v1.45.0
	gopkg.in/validator.v2 v2.0.1
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)
