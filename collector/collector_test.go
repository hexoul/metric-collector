package collector

import (
	"testing"
	"time"
)

func TestGenerateRequest(t *testing.T) {
	GenerateRequest()
	time.Sleep(10 * time.Second)
}
