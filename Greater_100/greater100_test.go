package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGreaterNumber(t *testing.T) {
    tests := []struct {
        a int
        want  bool
    }{
        {230, true},
        {101, true},
        {100, false},
        {29, false},
        {-9, false},
    }
    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            got := isGreater(tt.a)
            assert.Equal(t, tt.want, got)
        })
    }
}