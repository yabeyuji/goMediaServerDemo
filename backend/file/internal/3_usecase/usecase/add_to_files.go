package usecase

import "file/internal/4_domain/domain"

// AddtoFiles ...
func (uc *UseCase) AddtoFiles(files []domain.File, vid, name string) []domain.File {
	return uc.ToDomain.AddtoFiles(files, vid, name)
}
