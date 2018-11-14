package e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMsg(t *testing.T) {
	assert.Equal(t, "ok", GetMsg(200))
	assert.Equal(t, "Token error", GetMsg(20004))
}
