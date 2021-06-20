package usecase

import (
	"device/internal/4_domain/domain"
)

// GetRoomsString ...
func (uc *UseCase) GetRoomsString(rooms map[string]*domain.Room) (string, error) {
	raw, err := uc.ToDomain.JSONMarshal(rooms)
	if err != nil {
		myErr.Logging(err, rooms)
		return "", err
	}

	return string(raw), nil
}
