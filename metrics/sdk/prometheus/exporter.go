package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"sync"
)

var once sync.Once
var gRegistry *prometheus.Registry

func RegistryInstance() *prometheus.Registry {
	once.Do(func() {
		gRegistry = prometheus.NewRegistry()
	})
	return gRegistry
}

func RegisterMetric(metric *prometheus.GaugeVec) {
	RegistryInstance().Register(metric)
}

func NewGaugeVecMetric(name string, explanation string, labels []string) *prometheus.GaugeVec {
	return promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: explanation,
		},
		labels,
	)
}
