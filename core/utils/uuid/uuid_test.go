package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestNewUuid(t *testing.T) {
	uid := NewUuid()
	assert.NotEmpty(t, uid, "NewUuid() return empty")
}
