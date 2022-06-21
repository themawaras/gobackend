package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToke(t *testing.T) {

	result := NewToken("mario")

	assert.Equal(t, "mario", result.Username, "wrong")

}

func TestCreate(t *testing.T) {

	example := &Claims{
		Email: "toto",
	}

	result, _ := example.Create()

	assert.Equal(t, 159, len(result), "length result must be 180 char")
}
