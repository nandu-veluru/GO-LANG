package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInRange(t *testing.T) {
	tests := []struct {
		a    int
		want bool
	}{
		{50, true},
		{39, false},
		{-98, false},
		{0, false},
		{100, true},
		{-1, false},
		{52, true},
		{89, true},
		{math.MaxInt, false},
		{math.MinInt, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := inRange(tt.a)
			assert.Equal(t, tt.want, got)
		})
	}
}
