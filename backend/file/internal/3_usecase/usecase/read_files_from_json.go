package usecase

import "file/internal/4_domain/domain"

// ReadFilesFromJSON ...
func (uc *UseCase) ReadFilesFromJSON(filePath string) (*[]domain.File, error) {
	raw, err := uc.ToService.SvReadJSONFile(filePath)
	if err != nil {
		myErr.Logging(err, filePath)
		return nil, err
	}

	var files []domain.File
	err = uc.ToDomain.JSONUnmarshal(raw, &files)
	if err != nil {
		myErr.Logging(err, &files)
		return nil, err
	}

	return &files, nil
}
