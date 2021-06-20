package usecase

// SendIRData ...
func (uc *UseCase) SendIRData(room string, irKey string) error {
	err := uc.ToService.SvSendIRData(room, irKey)
	if err != nil {
		myErr.Logging(err, room, irKey)
		return err
	}

	return nil
}
