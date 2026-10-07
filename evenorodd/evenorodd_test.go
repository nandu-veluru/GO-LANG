package main

import "testing"
import "github.com/stretchr/testify/assert"


func TestAdd(t *testing.T) {
	actual := isEven(10)

	assert.Equal(t, true, actual) 
}
