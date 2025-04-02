package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	result := Soma(3, 4)
	expected := 7

	assert.Equal(t, expected, result, "Expected %d but got %d", expected, result)
}
