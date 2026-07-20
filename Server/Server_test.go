package main_test

import (
	main "ServerStateAPI"
	"net/http"
	"testing"
)

func TestGetIndex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main.GetIndex(tt.w, tt.r, nil)
		})
	}
}
