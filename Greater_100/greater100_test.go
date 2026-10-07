package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTrueForNumberGreaterThan100(t *testing.T) {
	randomNumber := rand.Intn(100000) + 101
	expected := true
	actual := isGreater(randomNumber)
	assert.Equal(t, expected, actual)
}

func TestShouldReturnFalseForNumberLesserThan100(t *testing.T) {
	randomNumber := rand.Intn(100)
	expected := false
	actual := isGreater(randomNumber)
	assert.Equal(t, expected, actual)
}
