package app

import (
	"media/internal/1_infrastructure/grpc/grpcin"
	"media/internal/1_infrastructure/grpc/grpcout"
	"media/internal/1_infrastructure/media/mediaapp"
	"media/internal/1_infrastructure/media/mediaorder"
	"media/internal/2_adapter/controller"

	"media/pkg/shared"
)

func init() {
	shared.CheckDirectory(shared.LogPath)
}

type (
	app struct {
		GrpcIn   *grpcin.GrpcIn
		MediaApp *mediaapp.MediaApp
	}
)

// NewApp ...
func NewApp() *app {
	a := &app{}
	grpcOut := grpcout.NewToGrpcOut()
	mediaOrder := mediaorder.NewToMediaOrder()
	ctrl := controller.NewController(grpcOut, mediaOrder)
	a.GrpcIn = grpcin.NewGrpcIn(ctrl)
	a.MediaApp = mediaapp.NewMediaApp(ctrl)

	return a
}

// Start ...
func (a *app) Start() {
	a.MediaApp.Start()
	a.GrpcIn.Controller.BootStrap()
	a.GrpcIn.Start()
}
