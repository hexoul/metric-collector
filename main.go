package main

import (
	"github.com/hexoul/metric-collector/collector"
	"github.com/hexoul/metric-collector/imon"
)

func init() {

}

func main() {
	collector.New(10, imon.MakeRequestContext)
}
