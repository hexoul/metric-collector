// Package collector gather data points of metric using generator
package collector

import (
	"net/http"

	"github.com/hexoul/metric-collector/generator"
)

// Collector gathers metric data
type Collector struct {
}

// New returns a collector
func New(rpm uint, consumer func(string) (*http.Request, error)) *Collector {
	metrics := []string{}
	contextChan := make(chan *http.Request)
	gather := func(metric string) {
		if req, err := consumer(metric); err != nil {
			contextChan <- req
		}
	}

	if g, err := generator.New(metrics, gather, rpm); err != nil {
		g.Start()
	}

	return &Collector{}
}
