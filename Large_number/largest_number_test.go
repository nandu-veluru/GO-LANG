package large_number

import (
	"fmt"
	"math"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{10, 20, 20},
		{20, 40, 40},
		{-9, -1, -1},
		{29, 29, 29},
		{0, 0, 0},
		{0, 26, 26},
		{-10, -87, -10},
		{-3, -9, -3},
		{-5, 5, 5},
		{math.MaxInt, 0, math.MaxInt},
		{math.MinInt, math.MaxInt, math.MaxInt},
		{math.MaxInt, math.MaxInt - 1, math.MaxInt},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := largestoftwonumbers(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Got %d, want %d", ans, tt.want)
			}
		})
	}
}
