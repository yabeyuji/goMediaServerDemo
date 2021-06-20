package usecase

// ParseKey ...
func (uc *UseCase) ParseKey(key string) error {
	err := uc.ToDomain.ParseKey(key)
	if err != nil {
		myErr.Logging(err, key)
		return err
	}

	return nil
}
