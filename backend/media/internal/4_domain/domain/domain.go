package domain

import (
	"errors"
	"strconv"

	"media/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("media", "domain")
}

type (
	domain struct{}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

func (domain *domain) ParseKey(key string) error {
	var err error

	switch key {
	case shared.DataKeyStatus:
	case shared.DataKeyProgress:
	case shared.DataKeyPlayList:
	default:
		err = errors.New("not found key")
	}

	if err != nil {
		myErr.Logging(err, key)
		return err
	}

	return nil
}

func (domain *domain) StringToFloat32(valueString string) (float32, error) {
	valueFloat64, err := strconv.ParseFloat(valueString, 32)
	if err != nil {
		myErr.Logging(err, valueString)
		return 0, err
	}
	valueFloat32 := float32(valueFloat64)

	return valueFloat32, nil
}
