package main

import(
	"net/http"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rw http.ResponseWriter
		r  *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StatusHandler(tt.rw, tt.r)
		})
	}
}
