package usecase

import (
	"file/pkg/shared"
)

// Move ...
func (uc *UseCase) Move(vid string) error {
	tempMP4 := uc.ToDomain.FilepathJoin(shared.TempFilePath, vid+".mp4")
	tempGIF := uc.ToDomain.FilepathJoin(shared.TempFilePath, vid+".gif")
	amineGIF := uc.ToDomain.FilepathJoin(shared.AnimeFilePath, vid+".gif")
	videoMP4 := uc.ToDomain.FilepathJoin(shared.VideoFilePath, vid+".mp4")

	command := []string{"mv", "-f", tempGIF, amineGIF}
	_, err := uc.ToService.SvExecCommand(command)
	if err != nil {
		myErr.Logging(err, command)
		return err
	}

	command = []string{"mv", "-f", tempMP4, videoMP4}
	_, err = uc.ToService.SvExecCommand(command)
	if err != nil {
		myErr.Logging(err, command)
		return err
	}

	return nil
}
