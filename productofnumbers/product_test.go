package main

import (
	"math"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{2, 3, 6},
		{-2, 4, -8},
		{-3, -3, 9},
		{12, 24, 288},
		{89078, 0, 0},
		{1, 999, 999},
		{math.MaxInt, 1, math.MaxInt},
		{math.MaxInt, 0, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := multiply(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("multiply(%d, %d) = %d; want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}
