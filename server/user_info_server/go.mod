module github.com/PsychologicalExperiment/backEnd/server/user_info_server

replace github.com/PsychologicalExperiment/backEnd/api/user_info_server => ../../api/user_info_server

go 1.16

require (
	github.com/PsychologicalExperiment/backEnd v0.0.0-20220612081424-2db32d97b7ef
	//github.com/PsychologicalExperiment/backEnd/api/user_info_server v0.0.0-20220522145057-7affc71739be
	github.com/PsychologicalExperiment/backEnd/util v0.0.0-20220410153153-365d8503d6fa
	github.com/go-redis/redis/v8 v8.11.5
	google.golang.org/grpc v1.45.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
	gopkg.in/yaml.v2 v2.4.0
)
