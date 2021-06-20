package grpcin

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"

	"ws/internal/2_adapter/controller"
	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "infrastructure:grpcin")
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
	log.Println("start GRPC ws ------------------------- ")
	lis, err := net.Listen("tcp", shared.GRPCPortWs)
	if err != nil {
		myErr.Logging(err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	cpb.RegisterSendContentServiceServer(s, &gin.Server)

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
	go s.Controller.SendContentToAgents(cc)

	return &cpb.SendContentResponse{Message: "ok"}, nil
}
