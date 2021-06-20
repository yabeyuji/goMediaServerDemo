package grpcin

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"
	"github.com/pborman/uuid"

	"file/internal/2_adapter/controller"
	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "infrastructure:grpcin")
}

// GrpcIn ...
type GrpcIn struct {
	Server
}

// Server ...
type Server struct {
	cpb.UnimplementedSendContentServiceServer
	cpb.UnimplementedReceiveContentServiceServer
	cpb.UnimplementedPostFileServer
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
	log.Println("start GRPC file ------------------------- ")
	lis, err := net.Listen("tcp", shared.GRPCPortFile)
	if err != nil {
		myErr.Logging(err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	cpb.RegisterSendContentServiceServer(s, &gin.Server)
	cpb.RegisterReceiveContentServiceServer(s, &gin.Server)
	cpb.RegisterPostFileServer(s, &gin.Server)

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

	switch cc.Key {
	case shared.Valid:
		go s.Controller.ToggleValid(cc)
	}

	return &cpb.SendContentResponse{Message: "ok"}, nil
}

// ReceiveContentRPC ...
func (s *Server) ReceiveContentRPC(ctx context.Context, in *cpb.ReceiveContentRequest) (*cpb.ReceiveContentResponse, error) {
	var err error
	response := &cpb.ReceiveContentResponse{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectFile,
		Key:    shared.DataKeyFiles,
	}

	response.Value, err = s.Controller.FileOperation(in.GetMessage())
	if err != nil {
		myErr.Logging(err)
		return nil, err
	}

	return response, nil
}

// PostFileRPC ...
func (s *Server) PostFileRPC(stream cpb.PostFile_PostFileRPCServer) error {
	var (
		name   string
		chunks []byte
		err    error
	)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			myErr.Logging(err)
			return err
		}

		if mt := resp.GetName(); mt != "" {
			name = mt
		}

		if chunk := resp.GetData(); chunk != nil {
			chunks = append(chunks, chunk...)
		}
	}

	if err := stream.SendAndClose(&cpb.ResponsePostFile{Message: "OK"}); err != nil {
		myErr.Logging(err)
		return err
	}

	vid := uuid.NewUUID().String()

	if err = s.Controller.UploadFile(name, vid, &chunks); err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
