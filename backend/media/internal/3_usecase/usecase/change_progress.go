package usecase

// ChangeProgress ...
func (uc *UseCase) ChangeProgress(value string) error {
	valueFloat32, err := uc.ToDomain.StringToFloat32(value)
	if err != nil {
		myErr.Logging(err, value)
		return err
	}

	uc.ToService.SvChangeVlcProgress(valueFloat32)

	return nil
}
