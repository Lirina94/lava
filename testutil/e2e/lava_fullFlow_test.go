package e2e

import (
	"testing"
	"time"
)

func TestLavaProtocol(t *testing.T) {
	t.Skip("skipping lava protocol test")
	// default timeout same as `go test`
	timeout := time.Minute * 10

	if deadline, ok := t.Deadline(); ok {
		timeout = time.Until(deadline).Round(10 * time.Second)
	}

	runProtocolE2E(timeout)
}

func TestLavaSDK(t *testing.T) {
	t.Skip("skipping lava sdk test")
	// default timeout same as `go test`
	timeout := time.Minute * 10

	if deadline, ok := t.Deadline(); ok {
		timeout = time.Until(deadline).Round(10 * time.Second)
	}

	runSDKE2E(timeout)
}
