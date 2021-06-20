package main

import (
	app "media/internal/1_infrastructure"
)

func main() {
	a := app.NewApp()
	a.Start()
}
