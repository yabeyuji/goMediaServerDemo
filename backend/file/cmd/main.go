package main

import (
	app "file/internal/1_infrastructure"
)

func main() {
	a := app.NewApp()
	a.Start()
}
