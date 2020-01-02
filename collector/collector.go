package collector

import (
	"fmt"

	"github.com/robfig/cron"
)

// GenerateRequest generates a request to collect
func GenerateRequest() {
	c := cron.New()
	a := 5
	c.AddFunc("@every 1s", func() { test(a) })
	c.Start()
}

func test(arg int) {
	fmt.Printf("called w/ %d\n", arg)
}
