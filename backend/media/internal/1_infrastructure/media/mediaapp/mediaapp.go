package mediaapp

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"

	"media/internal/1_infrastructure/media/mediachannel"
	"media/internal/2_adapter/controller"
	"media/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("media", "infrastructure:media")
}

// MediaApp ...
type MediaApp struct {
	MediaList  *vlc.MediaList
	ListPlayer *vlc.ListPlayer
	Player     *vlc.Player
	Quit       chan struct{}
	VlcStatus  string
	Controller *controller.Controller
}

// NewMediaApp ...
func NewMediaApp(ctrl *controller.Controller) *MediaApp {
	var err error

	md := &MediaApp{}
	md.Controller = ctrl

	if err := vlc.Init(); err != nil {
		myErr.Logging(err)
	}

	if md.Player, err = vlc.NewPlayer(); err != nil {
		myErr.Logging(err)
	}

	// Create a new list player.
	if md.ListPlayer, err = vlc.NewListPlayer(); err != nil {
		myErr.Logging(err)
	}

	if err = md.ListPlayer.SetPlayer(md.Player); err != nil {
		myErr.Logging(err)
	}

	// Create a new media list.
	if md.MediaList, err = vlc.NewMediaList(); err != nil {
		myErr.Logging(err)
	}

	// Set player media list.
	if err = md.ListPlayer.SetMediaList(md.MediaList); err != nil {
		myErr.Logging(err)
	}

	return md
}

// Start ...
func (md *MediaApp) Start() {
	md.VlcStatus = shared.VlcStop
	go md.WaitChannelOrder()
	go md.VlcUp()
}

// VlcUp ...
func (md *MediaApp) VlcUp() {
	defer func() {
		err := md.ListPlayer.Stop()
		if err != nil {
			myErr.Logging(err)
		}
		err = md.ListPlayer.Release()
		if err != nil {
			myErr.Logging(err)
		}
		err = md.MediaList.Release()
		if err != nil {
			myErr.Logging(err)
		}
	}()

	var err error
	err = md.Player.SetFullScreen(true)
	if err != nil {
		myErr.Logging(err)
	}

	err = md.ListPlayer.SetPlaybackMode(vlc.Loop)
	if err != nil {
		myErr.Logging(err)
	}

	// Retrieve player event manager.
	manager, err := md.Player.EventManager()
	if err != nil {
		myErr.Logging(err)
	}

	// Register the media end reached event with the event manager.
	md.Quit = make(chan struct{})

	events := []vlc.Event{
		vlc.MediaPlayerTimeChanged,
		vlc.MediaPlayerEndReached,
	}

	var eventIDs []vlc.EventID
	for _, event := range events {
		eventID, err := manager.Attach(event, md.EventControl, nil)
		if err != nil {
			myErr.Logging(err)
		}
		eventIDs = append(eventIDs, eventID)
	}

	// De-register attached events.
	defer func() {
		for _, eventID := range eventIDs {
			manager.Detach(eventID)
		}
	}()

	<-md.Quit
	return
}

// EventControl ...
func (md *MediaApp) EventControl(event vlc.Event, userData interface{}) {
	switch event {
	case vlc.MediaPlayerEndReached:
		log.Println("ListPlayer end reached")

	case vlc.MediaPlayerTimeChanged:
		stats, err := md.Player.MediaPosition()
		if err != nil {
			myErr.Logging(err)
		}
		md.Controller.SendProgress(stats)
	}
}

// WaitChannelOrder ...
func (md *MediaApp) WaitChannelOrder() {
	go func() {
		for {
			status := <-mediachannel.Status
			md.changeVlcStatus(status)
		}
	}()

	go func() {
		for {
			progress := <-mediachannel.Progress
			md.changeVlcProgress(progress)
		}
	}()

	go func() {
		for {
			playListString := <-mediachannel.PlayList

			md.swapPlayList(playListString)
		}
	}()
}
