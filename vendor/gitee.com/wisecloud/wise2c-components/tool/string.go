package tool

import (
	"strings"

	"github.com/satori/go.uuid"
)

func UUID() string {
	return uuid.NewV4().String()
}

func Trim(str string) (s string) {
	s = strings.Trim(str, " ")
	return
}
