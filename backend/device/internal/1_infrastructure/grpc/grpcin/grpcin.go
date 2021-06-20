package grpcin

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"

	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"

	"device/internal/2_adapter/controller"
	"device/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "infrastructure:grpcin")
}

// GrpcIn ...
type GrpcIn struct {
	Server
}

// Server ...
type Server struct {
	cpb.UnimplementedSendContentServiceServer
	cpb.UnimplementedReceiveContentServiceServer
	Controller *controller.Controller
}

// NewGrpcIn ...
func NewGrpcIn(ctrl *controller.Controller) *GrpcIn {
	gin := &GrpcIn{}
	sv := &Server{}
	sv.Controller = ctrl
	gin.Server = *sv

	return gin
}

// Start ...
func (gin *GrpcIn) Start() {
	log.Println("start GRPC device ------------------------- ")
	lis, err := net.Listen("tcp", shared.GRPCPortDevice)
	if err != nil {
		myErr.Logging(err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	cpb.RegisterSendContentServiceServer(s, &gin.Server)
	cpb.RegisterReceiveContentServiceServer(s, &gin.Server)

	if err := s.Serve(lis); err != nil {
		myErr.Logging(err)
		log.Fatalf("failed to serve: %v", err)
	}
}

// SendContentRPC ...
func (s *Server) SendContentRPC(ctx context.Context, in *cpb.SendContentRequest) (*cpb.SendContentResponse, error) {
	cc := &shared.CommonContent{
		Room:   in.GetRoom(),
		Object: in.GetObject(),
		Key:    in.GetKey(),
		Value:  in.GetValue(),
	}

	go s.Controller.SendToDevice(cc)

	return &cpb.SendContentResponse{Message: "ok"}, nil
}

// ReceiveContentRPC ...
func (s *Server) ReceiveContentRPC(ctx context.Context, in *cpb.ReceiveContentRequest) (*cpb.ReceiveContentResponse, error) {
	var err error
	var value string

	switch in.GetMessage() {
	case shared.TargetGetRooms:
		value = s.Controller.ReceiveContent(in.GetMessage())
	default:
		err = errors.New("not found function")
	}
	if err != nil {
		myErr.Logging(err, in.GetMessage())
		return nil, err
	}

	response := &cpb.ReceiveContentResponse{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectDevice,
		Key:    shared.DataKeyDevices,
		Value:  value,
	}

	return response, nil
}
