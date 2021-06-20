package wsapp

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// ServerInfo ...
	ServerInfo struct {
		Address string
		Port    string
	}
)

// Index ...
func (wd *WsApp) Index(address, port string) echo.HandlerFunc {
	return func(c echo.Context) error {
		serverInfo := &ServerInfo{
			Address: address,
			Port:    port,
		}
		data := struct{ *ServerInfo }{ServerInfo: serverInfo}

		return c.Render(http.StatusOK, "index", data)
	}
}
