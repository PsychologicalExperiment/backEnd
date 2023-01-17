package recovery

import (
	"runtime/debug"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
)

func init() {
	grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
		log.Errorf("debug: %s", string(debug.Stack()))
		return nil
	})	
}