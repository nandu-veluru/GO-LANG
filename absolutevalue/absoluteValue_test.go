package main
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAbsValue(t *testing.T) {
	tests := [] struct {
		a int
		want int
	}
	{
		{-12, 12},
		{5, 5},
		{-897, 897},
		{-17, 17},
		{0, 0},
		{39, 39},
		{-98, 98},
	}
for _,tt := range tests {
	t.Run("", func(t *testing.t) {
		got := absValue(tt.a)
		assert.Equal(tt, tt.want, got) 
	})

   }
}