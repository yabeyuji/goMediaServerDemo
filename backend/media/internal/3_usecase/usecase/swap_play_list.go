package usecase

// SwapPlayList ...
func (uc *UseCase) SwapPlayList(playListString string) {
	uc.ToService.SvSwapPlayList(playListString)
}
