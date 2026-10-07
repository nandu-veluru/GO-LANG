package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTrueForEvenNumber(t *testing.T) {
	randomNum := rand.Intn(math.MaxInt)
	if randomNum%2 != 0 {
		randomNum++
	}
	expected := true
	actual := isEvenOrOdd(randomNum)
	assert.Equal(t, expected, actual)
}

func TestShouldReturnFalseForOddNumber(t *testing.T) {
	randomNum := rand.Intn(math.MaxInt)
	if randomNum%2 == 0 {
		randomNum++
	}
	expected := false
	actual := isEvenOrOdd(randomNum)
	assert.Equal(t, expected, actual)
}
