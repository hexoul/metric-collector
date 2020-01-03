package generator

import (
	"testing"
	"time"
)

func TestConsume(t *testing.T) {
	metrics := []string{"a", "b", "c"}
	g, _ := New(metrics, func(param string) {
		t.Log("Consumed", param)
	}, 2)
	g.Start()
	time.Sleep(30 * time.Second)

	// Expected output: a, b, c, a, b, c, a, b, ...
	// The order between a, b and c could be different.
}
