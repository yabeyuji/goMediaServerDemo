package main

import (
	app "ws/internal/1_infrastructure"
)

func main() {
	a := app.NewApp()
	a.Start()
}
