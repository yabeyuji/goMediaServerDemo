package app

import (
	"file/internal/1_infrastructure/fileapp"
	"file/internal/1_infrastructure/grpc/grpcin"
	"file/internal/1_infrastructure/grpc/grpcout"
	"file/internal/2_adapter/controller"

	"file/pkg/shared"
)

type (
	app struct {
		GrpcIn *grpcin.GrpcIn
	}
)

// NewApp ...
func NewApp() *app {
	a := &app{}
	grpcOut := grpcout.NewToGrpcOut()
	file := fileapp.NewToFileApp()
	ctrl := controller.NewController(grpcOut, file)
	a.GrpcIn = grpcin.NewGrpcIn(ctrl)
	return a
}

// Start ...
func (a *app) Start() {
	a.GrpcIn.Controller.BootStrap()
	a.GrpcIn.Start()
}

func init() {
	shared.CheckDirectory(shared.FileParePath)
	shared.CheckDirectory(shared.VideoFilePath)
	shared.CheckDirectory(shared.TempFilePath)
	shared.CheckDirectory(shared.AnimeFilePath)
	shared.CheckDirectory(shared.LogPath)
}
