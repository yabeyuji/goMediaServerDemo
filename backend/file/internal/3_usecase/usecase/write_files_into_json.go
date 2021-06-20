package usecase

import (
	"os"

	"file/internal/4_domain/domain"
	"file/pkg/shared"
)

// WriteFilesIntoJSON ...
func (uc *UseCase) WriteFilesIntoJSON(files *[]domain.File) error {
	raw, err := uc.ToDomain.JSONMarshalIndent(files)
	if err != nil {
		myErr.Logging(err, files)
		return err
	}

	err = uc.ToService.SvWriteJSONFile(shared.JSONPath, []byte(raw), os.ModePerm)
	if err != nil {
		myErr.Logging(err, shared.JSONPath, []byte(raw), os.ModePerm)
		return err
	}

	return nil
}
