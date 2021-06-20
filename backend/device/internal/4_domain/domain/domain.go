package domain

import (
	"encoding/json"
	"strconv"

	"device/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "domain")
}

type (
	domain struct{}

	// Room ...
	Room struct {
		Aircon *Aircon
		Light  *Light
	}

	// Aircon ...
	Aircon struct {
		Status          string
		WarmTemperature int
		CoolTemperature int
	}

	// Light ...
	Light struct {
		Status string
	}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

// GetDefaltValue ...
func (domain *domain) GetDefaltValue() *Room {
	return &Room{
		Light: &Light{
			Status: shared.DefaultLightStatus,
		},
		Aircon: &Aircon{
			Status:          shared.DefaultAirconStatus,
			WarmTemperature: shared.DefaultAirconWarmTemperature,
			CoolTemperature: shared.DefaultAirconCoolTemperature,
		},
	}
}

// GetBroadcastStatus ...
func (domain *domain) GetBroadcastStatus(object string) bool {
	var isBroadcastStatus bool

	switch object {
	case shared.DataObjectLight, shared.DataObjectAircon:
		isBroadcastStatus = true
	}

	return isBroadcastStatus
}

// JSONMarshal ...
func (domain *domain) JSONMarshal(rooms map[string]*Room) ([]byte, error) {
	return json.Marshal(&rooms)
}

// ChangeDeviceStatus ...i
func (domain *domain) ChangeDeviceStatus(cc *shared.CommonContent, rooms map[string]*Room) map[string]*Room {

	switch cc.Object {
	case shared.DataObjectLight:
		rooms[cc.Room].Light.Status = cc.Value

	case shared.DataObjectAircon:

		switch cc.Key {
		case shared.DataKeyStatus:
			rooms[cc.Room].Aircon.Status = cc.Value

		case shared.DataKeyWarmTemperature:
			templature, err := domain.StrConvAtoi(cc.Value)
			if err != nil {
				myErr.Logging(err, cc.Value)
			}
			rooms[cc.Room].Aircon.WarmTemperature = templature

		case shared.DataKeyCoolTemperature:
			templature, err := domain.StrConvAtoi(cc.Value)
			if err != nil {
				myErr.Logging(err, cc.Value)
			}
			rooms[cc.Room].Aircon.CoolTemperature = templature
		}

	}

	return rooms
}

// StrConvAtoi ...
func (domain *domain) StrConvAtoi(valueString string) (int, error) {
	valueInt, err := strconv.Atoi(valueString)
	if err != nil {
		myErr.Logging(err, valueString)
		return 0, err
	}
	return valueInt, nil
}
