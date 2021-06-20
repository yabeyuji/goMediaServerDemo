package usecase

import (
	"file/pkg/shared"
)

// Convert ...
func (uc *UseCase) Convert(vid string) error {
	tempMP4 := uc.ToDomain.FilepathJoin(shared.TempFilePath, vid+".mp4")
	tempGIF := uc.ToDomain.FilepathJoin(shared.TempFilePath, vid+".gif")

	command := []string{
		"ffmpeg",
		"-i", tempMP4,
		"-t", "20",
		"-vf", "scale=320:-1",
		"-r", "10",
		"-y", tempGIF,
	}
	_, err := uc.ToService.SvExecCommand(command)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
