package usecase

import "file/internal/4_domain/domain"

// GetValidFiles ...
func (uc *UseCase) GetValidFiles(files *[]domain.File) string {
	filesString := uc.ToDomain.GetValidFiles(files)

	return filesString
}
