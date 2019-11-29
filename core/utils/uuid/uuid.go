package uuid

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewUuid() string {
	id := uuid.NewV4().String()

	return strings.ReplaceAll(id, "-", "")
}
