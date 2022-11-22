package utils

import (
	"github.com/stretchr/testify/assert"
	"boilerplate/env"
	"testing"
)

func Test_maybePrependLocalhost(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		addr     string
		expected string
	}{
		{
			name:     "Prod - port only",
			env:      env.Prod,
			addr:     ":3000",
			expected: ":3000",
		},
		{
			name:     "Prod - host and port",
			env:      env.Prod,
			addr:     "127.0.0.1:3000",
			expected: "127.0.0.1:3000",
		},
		{
			name:     "Dev - no port",
			env:      env.Dev,
			addr:     "3000",
			expected: "3000",
		},
		{
			name:     "Dev - host and port",
			env:      env.Dev,
			addr:     "0.0.0.0:3000",
			expected: "0.0.0.0:3000",
		},
		{
			name:     "Prod - empty",
			env:      env.Prod,
			addr:     "",
			expected: "",
		},
		{
			name:     "Dev - empty",
			env:      env.Dev,
			addr:     "",
			expected: "127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env.SetForTesting(t, tt.env)
			assert.Equal(t, tt.expected, MaybePrependLocalhost(tt.addr))
		})
	}
}
