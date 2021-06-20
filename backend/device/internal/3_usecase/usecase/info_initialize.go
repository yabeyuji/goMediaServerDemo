package usecase

import (
	"device/internal/4_domain/domain"
	"device/pkg/shared"
)

// InfoInitialize ...
func (uc *UseCase) InfoInitialize() map[string]*domain.Room {
	rooms := make(map[string]*domain.Room)
	rooms[shared.DataRoomLiving] = uc.ToDomain.GetDefaltValue()
	rooms[shared.DataRoomBed] = uc.ToDomain.GetDefaltValue()

	return rooms
}
