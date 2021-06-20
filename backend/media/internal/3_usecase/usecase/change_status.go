package usecase

// ChangeStatus ...
func (uc *UseCase) ChangeStatus(status string) {
	uc.ToService.SvChangeStatus(status)
}
