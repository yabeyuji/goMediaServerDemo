package main

import (
	app "device/internal/1_infrastructure"
)

func main() {
	a := app.NewApp()
	a.Start()
}
