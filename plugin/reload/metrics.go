package reload

import (
	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics for the reload plugin
var (
	// failedCount is the counter of the number of failed reload attempts.
	failedCount = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "reload",
		Name:      "failed_total",
		Help:      "Counter of the number of failed reload attempts.",
	})
	// reloadInfo is record the hash value during reload.
	reloadInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: plugin.Namespace,
		Subsystem: "reload",
		Name:      "version_info",
		Help:      "A metric with a constant '1' value labeled by hash, and value which type of hash generated.",
	}, []string{"hash", "value"})
	// lastReloadSuccess records 1 if the last configuration reload was successful, 0 if it failed.
	lastReloadSuccess = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: plugin.Namespace,
		Subsystem: "reload",
		Name:      "config_last_reload_success",
		Help:      "The last configuration reload succeeded.",
	})
)

func init() {
	lastReloadSuccess.Set(1)
}
