package generator

import (
	"github.com/robfig/cron"
)

// Generator keeps a cron to make requests periodically
type Generator struct {
	cron    *cron.Cron
	metrics []string
	provide chan string
	rpm     uint
}

// New returns a new request generator collect.
// consume will be called with each metrics, thereby consume have need to get string as parameter.
// rpm means Requests Per Minute referred by consume. It should be bigger than zero
// and smaller than the length of metrics.
func New(
	metrics []string,
	consume func(string),
	rpm uint,
) (*Generator, error) {
	// verify parameters

	c := cron.New()
	provide := make(chan string, len(metrics))

	_, err := c.AddFunc("@every 20s", func() { feed(metrics, provide) })
	if err != nil {
		return nil, err
	}

	_, err = c.AddFunc("@every 1m", func() {
		for i := uint(0); i < rpm; i++ {
			consume(<-provide)
		}
	})
	if err != nil {
		return nil, err
	}

	return &Generator{
		cron:    c,
		metrics: metrics,
		rpm:     rpm,
	}, nil
}

func feed(metrics []string, provide chan string) {
	if len(provide) >= len(metrics) {
		return
	}

	for _, metric := range metrics {
		go func(metric string) { provide <- metric }(metric)
	}
}

// Start the generator
func (g *Generator) Start() {
	g.cron.Start()
}
