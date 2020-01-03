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
	time.Sleep(150 * time.Second)

	// Expected output: a, b, c, a
	// The output order and last item could be different.
}
