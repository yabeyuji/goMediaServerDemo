package usecase

// GetBroadcastStatus ...
func (uc *UseCase) GetBroadcastStatus(object string) bool {
	return uc.ToDomain.GetBroadcastStatus(object)
}
