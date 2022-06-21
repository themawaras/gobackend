package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {

	result := NewToken("vina")

	assert.Equal(t, "vina", result.Username, "wrong")

}
