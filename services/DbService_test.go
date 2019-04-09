package services

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup

	code := m.Run()
	// teardown

	os.Exit(code)
}

func TestQueryBrands(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "testdb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueryBrands()
		})
	}
}
