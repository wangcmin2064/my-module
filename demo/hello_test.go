package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	msg := Hello()
	assert.Equal(t, "hello world!111111333333", msg)
}
