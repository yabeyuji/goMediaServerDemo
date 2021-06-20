package usecase

import (
	"device/internal/4_domain/domain"
	"device/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "usecase")
}

type (
	// UseCase ...
	UseCase struct {
		ToDomain  ToDomain
		ToService ToService
	}

	// ToService ...
	ToService interface {
		SvSendIRData(room string, irKey string) error
		SvSendContent(address string, cc *shared.CommonContent) (string, error)
	}

	// ToDomain ...
	ToDomain interface {
		GetDefaltValue() *domain.Room
		GetBroadcastStatus(object string) bool
		JSONMarshal(rooms map[string]*domain.Room) ([]byte, error)
		StrConvAtoi(valueString string) (int, error)
		ChangeDeviceStatus(cc *shared.CommonContent, rooms map[string]*domain.Room) map[string]*domain.Room
	}
)
