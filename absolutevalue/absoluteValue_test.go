package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsValue(t *testing.T) {
	tests := []struct {
		a    int
		want int
	}{
		{-12, 12},
		{5, 5},
		{-897, 897},
		{-17, 17},
		{0, 0},
		{39, 39},
		{-98, 98},
		{2, 2},
		{-90, 90},
		{7865, 7865},
		{-878, 878},
		{-98, 98},
		{-9, 9},
		{-1058238, 1058238},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := absValue(tt.a)
			assert.Equal(t, tt.want, got)
		})

	}
}
