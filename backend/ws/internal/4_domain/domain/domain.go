package domain

import (
	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "domain")
}

type (
	domain struct{}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}
