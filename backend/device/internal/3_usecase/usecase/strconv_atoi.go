package usecase

// StrConvAtoi ...
func (uc *UseCase) StrConvAtoi(valueString string) (int, error) {
	valueInt, err := uc.ToDomain.StrConvAtoi(valueString)
	if err != nil {
		myErr.Logging(err, valueString)
		return 0, err
	}

	return valueInt, nil
}
