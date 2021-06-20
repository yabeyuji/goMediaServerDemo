package shared

import (
	"bytes"
	"context"
	"io"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"

	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"
)

// ReceiveContent ...
func ReceiveContent(address, message string) (*cpb.ReceiveContentResponse, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cpbconn := cpb.NewReceiveContentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cpbconn.ReceiveContentRPC(ctx, &cpb.ReceiveContentRequest{Message: message})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SendContent ...
func SendContent(address string, request *cpb.SendContentRequest) (*cpb.SendContentResponse, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cpbconn := cpb.NewSendContentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cpbconn.SendContentRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// PostFile ...
func PostFile(address, filename string, fileBody *bytes.Buffer) (*cpb.ResponsePostFile, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cpbconn := cpb.NewPostFileClient(conn)

	stream, err := cpbconn.PostFileRPC(context.Background())
	if err != nil {
		return nil, err
	}

	name := strings.Split(filename, ".")[0]
	stream.Send(&cpb.RequestPostFile{Value: &cpb.RequestPostFile_Name{Name: name}})

	data := make([]byte, 1024)
	for {
		_, err := fileBody.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Could not send file: %v.", err)
			break
		}
		stream.Send(&cpb.RequestPostFile{Value: &cpb.RequestPostFile_Data{Data: data}})
	}

	return stream.CloseAndRecv()
}
