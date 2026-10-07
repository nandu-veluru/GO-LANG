package large_number
import (
"fmt"
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


