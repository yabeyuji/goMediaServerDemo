package deviceapp

import (
	"fmt"
	"net/http"
	"net/url"

	// url "net/url"

	"device/internal/2_adapter/service"
	"device/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "infrastructure:device")
}

type device struct {
	irData map[string]string
}

// NewToDevice ...
func NewToDevice() service.ToDevice {
	dv := new(device)
	dv.irData = getIRData()
	return dv
}

// IsSendIRData ...
func (dv *device) IsSendIRData(room string, irKey string) error {
	servers := map[string]string{
		shared.DataRoomLiving: "192.168.8.210",
		shared.DataRoomBed:    "192.168.8.211",
		// shared.DataRoomLiving: "serverLiving",
		// shared.DataRoomBed:    "serverBed",
	}

	rawuri := fmt.Sprint("http://", servers[room], ":4001/", dv.irData[irKey])
	u, err := url.Parse(rawuri)
	if err != nil {
		myErr.Logging(err, rawuri)
		return err
	}

	_, err = http.Get(u.String())
	if err != nil {
		myErr.Logging(err, u.String())
		return err
	}

	return nil
}
