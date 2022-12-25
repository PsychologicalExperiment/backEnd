module github.com/PsychologicalExperiment/backEnd/server/user_info_server

replace github.com/PsychologicalExperiment/backEnd/api/user_info_server => ../../api/user_info_server

go 1.16


require (
	github.com/PsychologicalExperiment/backEnd v0.0.0-20221225100123-0ae180cc4328
	github.com/PsychologicalExperiment/backEnd/api/user_info_server v0.0.0-00010101000000-000000000000
	github.com/PsychologicalExperiment/backEnd/util v0.0.0-20221225100123-0ae180cc4328
	github.com/go-redis/redis/v8 v8.11.0
	google.golang.org/grpc v1.51.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)
