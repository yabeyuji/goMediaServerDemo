package wsapp

import (
	"io"
	"log"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"ws/internal/2_adapter/controller"
	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "infrastructure:ws")
}

type (
	// Template ...
	Template struct {
		templates *template.Template
	}

	// WsApp ...
	WsApp struct {
		EchoEcho   *echo.Echo
		Agents     map[string]*Agent
		Mutex      sync.RWMutex
		Member     []string
		Controller *controller.Controller
	}

	// Agent ...
	Agent struct {
		ID     string
		Socket *websocket.Conn
	}
)

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewWsApp ...
func NewWsApp(ctrl *controller.Controller) *WsApp {
	wd := &WsApp{}
	wd.EchoEcho = NewEcho()
	wd.Controller = ctrl
	wd.Agents = make(map[string]*Agent)

	return wd
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob(shared.IndexFilePath)),
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}__${status}__${method}__${uri}\n",
	}))
	e.Use(middleware.Recover())

	e.Static("/", "public")

	return e
}

// Start ...
func (wd *WsApp) Start(address, port string) {
	log.Println("--------------------------- ")
	log.Println("http://" + address + ":" + port)
	log.Println("--------------------------- ")

	go wd.SendToAgents()

	wd.EchoEcho.Static("/public", shared.PublicPath)
	wd.EchoEcho.GET("/", wd.Index(address, port))
	wd.EchoEcho.GET("/ws", wd.WebSocket)
	wd.EchoEcho.POST("/file_upload", wd.FileUpload)
	wd.EchoEcho.Logger.Fatal(wd.EchoEcho.Start(":" + port))
}
