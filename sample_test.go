package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	result := HelloName("world")
	assert.Equal(t, "Hello world", result, "must be Hello world")
}
