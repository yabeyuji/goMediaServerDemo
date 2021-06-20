package usecase

import (
	"device/internal/4_domain/domain"
	"device/pkg/shared"
)

// ChangeDeviceStatus ...
func (uc *UseCase) ChangeDeviceStatus(cc *shared.CommonContent, rooms map[string]*domain.Room) map[string]*domain.Room {
	return uc.ToDomain.ChangeDeviceStatus(cc, rooms)
}
