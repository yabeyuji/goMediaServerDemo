package app

import (
	"device/internal/1_infrastructure/deviceapp"
	"device/internal/1_infrastructure/grpc/grpcin"
	"device/internal/1_infrastructure/grpc/grpcout"
	"device/internal/2_adapter/controller"
	"device/pkg/shared"
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
	device := deviceapp.NewToDevice()
	ctrl := controller.NewController(grpcOut, device)
	a.GrpcIn = grpcin.NewGrpcIn(ctrl)
	return a
}

// Start ...
func (a *app) Start() {
	a.GrpcIn.Controller.BootStrap()
	a.GrpcIn.Start()
}

func init() {
	shared.CheckDirectory(shared.LogPath)
}
