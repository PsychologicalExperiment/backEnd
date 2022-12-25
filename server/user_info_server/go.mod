module github.com/PsychologicalExperiment/backEnd/server/user_info_server

replace github.com/PsychologicalExperiment/backEnd/api/user_info_server => ../../api/user_info_server

go 1.16

//replace (
//	"github.com/PsychologicalExperiment/backEnd/api" => "../../api"
//)

require (
	github.com/go-redis/redis/v8 v8.11.0
	google.golang.org/grpc v1.51.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)
