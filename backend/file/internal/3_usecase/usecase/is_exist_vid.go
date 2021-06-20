package usecase

import "file/internal/4_domain/domain"

// IsExistVid ...
func (uc *UseCase) IsExistVid(files *[]domain.File, vid string) bool {
	ok := uc.ToDomain.IsExistVid(files, vid)

	return ok
}
