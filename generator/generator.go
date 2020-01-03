// Package generator provides a periodic request generation
package generator

import (
	"github.com/robfig/cron"
)

// Generator keeps a cron to make requests periodically
type Generator struct {
	cron        *cron.Cron
	metrics     []string
	provide     chan string
	lastProvide int
	rpm         uint
}

// New returns a new request generator collect.
// consumer will be called with each metrics, thereby consumer have need to get string as parameter.
// rpm means Requests Per Minute referred by consumer. It should be bigger than zero.
func New(
	metrics []string,
	consumer func(string),
	rpm uint,
) (*Generator, error) {
	c := cron.New()
	p := make(chan string, 1000)
	g := &Generator{
		cron:        c,
		metrics:     metrics,
		provide:     p,
		lastProvide: 0,
		rpm:         rpm,
	}

	_, err := c.AddFunc("@every 10s", func() { g.feed() })
	if err != nil {
		return nil, err
	}

	_, err = c.AddFunc("@every 1m", func() { g.consume(consumer) })
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Generator) feed() {
	diff := cap(g.provide) - len(g.provide)
	if diff == 0 {
		return
	}

	iter := diff
	if len(g.metrics) < iter {
		iter = len(g.metrics)
	}
	for i := 0; i < iter; i++ {
		idx := (i + g.lastProvide) % len(g.metrics)
		go func(metric string) { g.provide <- metric }(g.metrics[idx])
	}
	g.lastProvide += iter
}

func (g *Generator) consume(consumer func(string)) {
	for i := uint(0); i < g.rpm; i++ {
		consumer(<-g.provide)
	}
}

// Start the generator
func (g *Generator) Start() {
	g.cron.Start()
}

// Stop the generator
func (g *Generator) Stop() {
	g.cron.Stop()
}
