package usecase

import "file/internal/4_domain/domain"

// ToggleFiles ...
func (uc *UseCase) ToggleFiles(files []domain.File, vid string) []domain.File {
	return uc.ToDomain.ToggleFiles(files, vid)
}
