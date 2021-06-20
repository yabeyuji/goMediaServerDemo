package usecase

import (
	"file/pkg/shared"
)

// UploadFile ...
func (uc *UseCase) UploadFile(vid string, chunks *[]byte) error {
	tempMP4 := uc.ToDomain.FilepathJoin(shared.TempFilePath, vid+".mp4")

	if err := uc.ToService.SvUploadFile(tempMP4, chunks); err != nil {
		return err
	}

	return nil
}
