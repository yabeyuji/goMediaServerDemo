package mediaorder

import (
	"media/internal/1_infrastructure/media/mediachannel"
	"media/internal/2_adapter/service"
)

type mediaOrder struct{}

// NewToMediaOrder ...
func NewToMediaOrder() service.ToMediaOrder {
	mo := new(mediaOrder)

	return mo
}

// IsChangeStatus ...
func (order *mediaOrder) IsChangeStatus(status string) {
	mediachannel.Status <- status

	return
}

// IsChangeVlcProgress ...
func (order *mediaOrder) IsChangeVlcProgress(value float32) {
	mediachannel.Progress <- value

	return
}

// IsSwapPlayList ...
func (order *mediaOrder) IsSwapPlayList(playListString string) {
	mediachannel.PlayList <- playListString

	return
}
