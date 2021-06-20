package usecase

import (
	"time"

	"file/pkg/shared"
)

// IsExistFiles ...
func (uc *UseCase) IsExistFiles(vid string) bool {
	mp4Path := uc.ToDomain.FilepathJoin(shared.VideoFilePath, vid+".mp4")
	gifPath := uc.ToDomain.FilepathJoin(shared.AnimeFilePath, vid+".gif")

	cnt := 0
	for {
		mp4ok := uc.ToDomain.IsExistFilePath(mp4Path)
		gifok := uc.ToDomain.IsExistFilePath(gifPath)

		if mp4ok && gifok {
			return true
		}

		if cnt == 10 {
			break
		}

		time.Sleep(time.Second * 1)
		cnt++
	}

	return false
}
