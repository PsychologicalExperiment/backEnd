module github.com/PsychologicalExperiment/backEnd/client/experiment_client

replace github.com/PsychologicalExperiment/backEnd/api/experiment_server => ../../api/experiment_server

go 1.16

require (
	github.com/PsychologicalExperiment/backEnd v0.0.0-20220410153153-365d8503d6fa // indirect
	github.com/PsychologicalExperiment/backEnd/api/experiment_server v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.46.0
)
