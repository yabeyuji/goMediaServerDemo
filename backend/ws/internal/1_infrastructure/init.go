package app

import (
	"ws/internal/1_infrastructure/grpc/grpcin"
	"ws/internal/1_infrastructure/grpc/grpcout"
	"ws/internal/1_infrastructure/network"
	"ws/internal/1_infrastructure/ws/wsapp"
	"ws/internal/1_infrastructure/ws/wsorder"
	"ws/internal/2_adapter/controller"

	"ws/pkg/shared"
)

func init() {
	shared.CheckDirectory(shared.LogPath)
}

type (
	app struct {
		GrpcIn *grpcin.GrpcIn
		WsApp  *wsapp.WsApp
	}
)

// NewApp ...
func NewApp() *app {
	a := &app{}

	grpcOut := grpcout.NewToGrpcOut()
	wsOrder := wsorder.NewToWsOrder()
	ctrl := controller.NewController(grpcOut, wsOrder)
	a.GrpcIn = grpcin.NewGrpcIn(ctrl)
	a.WsApp = wsapp.NewWsApp(ctrl)

	return a
}

// Start ...
func (a *app) Start() {
	go a.WsApp.Start(network.GetSelfAddress(), shared.EchoPort)
	a.GrpcIn.Start()
}
