package mon

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/naming"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	ip, err := naming.GetExternalIP()
	if err != nil {
		log.Infof("get ip error")
		os.Exit(2)
	}
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf("%s:%d", ip, config.Config.Monitor.Port),
	}
	go func() {
		log.Infof("start promethues server")
		if err := httpServer.ListenAndServe(); err != nil {
			os.Exit(2)
		}
	}()
}
