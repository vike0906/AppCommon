package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

//the UUID based on random numbers
func UUID() string {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	return strings.Replace(uuidStr, "-", "", -1)
}
