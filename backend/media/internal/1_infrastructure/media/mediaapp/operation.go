package mediaapp

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"media/pkg/shared"
)

func (md *MediaApp) swapPlayList(playListString string) {
	for {
		cnt, err := md.MediaList.Count()
		if err != nil {
			myErr.Logging(err)
		}
		if cnt == 0 {
			break
		}
		for i := 0; i <= cnt; i++ {
			var z uint = uint(i)
			err := md.MediaList.RemoveMediaAtIndex(z)
			if err != nil {
				myErr.Logging(err)
			}

		}
		time.Sleep(time.Millisecond * 10)
	}

	for {
		if f, err := os.Stat(shared.VideoFilePath); !os.IsNotExist(err) || f.IsDir() {
			break
		}
		time.Sleep(time.Second * 1)
	}

	playLists := strings.Split(playListString, ",")

	for _, validFile := range playLists {
		if validFile != "" {
			md.addVideo(validFile)
		}
	}
}

func (md *MediaApp) addVideo(validFile string) {
	filePath := filepath.Join(shared.VideoFilePath, validFile+".mp4")
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		err := md.MediaList.AddMediaFromPath(filePath)
		if err != nil {
			myErr.Logging(err, filePath)
			log.Fatal(err)
		}
	}
}

func (md *MediaApp) changeVlcProgress(progress float32) {
	if md.VlcStatus != shared.VlcStop {
		err := md.Player.SetMediaPosition(progress)
		if err != nil {
			myErr.Logging(err, progress)
		}
	}
}

func (md *MediaApp) changeVlcStatus(targetStatus string) {
	var err error
	newStatus := md.VlcStatus

	switch {
	case targetStatus == shared.VlcPlay && md.VlcStatus != shared.VlcPlay:
		newStatus = shared.VlcPlay
		err = md.ListPlayer.Play()

	case targetStatus == shared.VlcStop && (md.VlcStatus == shared.VlcPause || md.VlcStatus == shared.VlcPlay):
		newStatus = shared.VlcStop
		err = md.ListPlayer.Stop()

	case targetStatus == shared.VlcPause && (md.VlcStatus == shared.VlcPause || md.VlcStatus == shared.VlcPlay):
		switch md.VlcStatus {
		case shared.VlcPause:
			newStatus = shared.VlcPlay
			err = md.ListPlayer.Play()
		case shared.VlcPlay:
			newStatus = shared.VlcPause
			err = md.ListPlayer.SetPause(true)
		}

	case targetStatus == shared.VlcPauseOn:
		newStatus = shared.VlcPause
		err = md.ListPlayer.SetPause(true)

	case targetStatus == shared.VlcPauseOff:
		newStatus = shared.VlcPlay
		err = md.ListPlayer.SetPause(false)

	case targetStatus == shared.VlcPlayNext:
		newStatus = shared.VlcPlay
		err = md.ListPlayer.PlayNext()

	case targetStatus == shared.VlcPlayPrevious:
		newStatus = shared.VlcPlay
		err = md.ListPlayer.PlayPrevious()

	default:
		err = errors.New("not found status")
	}

	if err != nil {
		myErr.Logging(err, targetStatus, newStatus)
	}

	md.VlcStatus = newStatus

	return
}
