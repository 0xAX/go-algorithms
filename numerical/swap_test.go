package numerical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	x := 5
	y := 6

	expectedX := 6
	expectedY := 5

	swap(&x, &y)
	assert.Equal(t, x, expectedX, "value should be equal")
	assert.Equal(t, y, expectedY, "value should be equal")
}
