package recovery

import (
	"runtime/debug"

	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func init() {
	grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
		log.Errorf("panic stack: %s", string(debug.Stack()))
		return nil
	})
}
