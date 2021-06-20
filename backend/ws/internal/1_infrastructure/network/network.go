package network

import (
	"net"
	"strings"
	"time"
)

func GetSelfAddress() (address string) {
	for {
		netAddrList, _ := net.InterfaceAddrs()
		for _, netAddr := range netAddrList {
			na, ok := netAddr.(*net.IPNet)
			if ok && strings.HasPrefix(na.IP.String(), "192.168.") {
				address = na.IP.String()
				break
			}
		}
		if address != "" {
			break
		}
		time.Sleep(time.Second * 1)
	}

	return address
}
