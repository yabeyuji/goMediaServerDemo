package shared

import (
	"os"
	"path/filepath"
)

const (
	GRPCPortWs        = ":40001"
	GRPCPortMedia     = ":40011"
	GRPCPortFile      = ":40021"
	GRPCPortDevice    = ":40031"
	GRPCAddressMedia  = "localhost" + GRPCPortMedia
	GRPCAddressWs     = "localhost" + GRPCPortWs
	GRPCAddressFile   = "localhost" + GRPCPortFile
	GRPCAddressDevice = "localhost" + GRPCPortDevice
	EchoPort          = "4000"
)

const (
	VlcPlay         = "vlcPlay"
	VlcStop         = "vlcStop"
	VlcPause        = "vlcPause"
	VlcPauseOn      = "vlcPauseOn"
	VlcPauseOff     = "vlcPauseOff"
	VlcPlayNext     = "vlcPlayNext"
	VlcPlayPrevious = "vlcPlayPrevious"
)

const (
	DataRoomCommon = "common"
	DataRoomLiving = "living"
	DataRoomBed    = "bed"

	DataObjectVlc    = "vlc"
	DataObjectClient = "client"
	DataObjectFile   = "file"
	DataObjectDevice = "device"

	DataObjectLight     = "light"
	DataObjectProjector = "projector"
	DataObjectAircon    = "aircon"
	DataObjectMonitor   = "monitor"
	DataObjectTV        = "tv"

	DataKeyStatus   = "status"
	DataKeyPlayList = "playList"
	DataKeyEnter    = "enter"
	DataKeyMember   = "member"
	DataKeyLeave    = "leave"
	DataKeyProgress = "progress"
	DataKeyFiles    = "files"
	DataKeyDevices  = "devices"

	DataKeyWarmTemperature = "warmTemperature"
	DataKeyCoolTemperature = "coolTemperature"
)

const (
	Valid = "valid"
)

var (
	currentPath, _ = os.Getwd()
	rootPath       = filepath.Dir(currentPath)
	storagePath    = filepath.Join(rootPath, "storage")
	LogPath        = filepath.Join(storagePath, "log")
	FileParePath   = filepath.Join(storagePath, "file")
	JSONPath       = filepath.Join(FileParePath, "db.json")
	VideoFilePath  = filepath.Join(FileParePath, "video")
	TempFilePath   = filepath.Join(FileParePath, "temp")
	AnimeFilePath  = filepath.Join(FileParePath, "anime")
	PublicPath     = filepath.Join(rootPath, "web", "public")
	IndexFilePath  = filepath.Join(PublicPath, "*.html")
)

const (
	DefaultLightStatus  = "lightPower"
	DefaultAirconStatus = "airconStop"
)

const (
	DefaultAirconWarmTemperature = 23
	DefaultAirconCoolTemperature = 18
)

const (
	TargetGetValidFiles = "getValidFiles"
	TargetGetFiles      = "getFiles"
	TargetGetRooms      = "getRooms"
)
