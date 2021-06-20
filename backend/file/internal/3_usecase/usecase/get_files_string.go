package usecase

import "file/internal/4_domain/domain"

// GetFilesString ...
func (uc *UseCase) GetFilesString(files *[]domain.File) (string, error) {
	raw, err := uc.ToDomain.JSONMarshal(files)
	if err != nil {
		myErr.Logging(err, files)
		return "", err
	}

	return string(raw), nil
}
