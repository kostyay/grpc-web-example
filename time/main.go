package main

import (
	"context"
	"flag"
	"log/slog"
	"net"
	"os"
	"time"

	pb "github.com/kostyay/grpc-web-example/time/goclient/api/time/v1"
	"google.golang.org/grpc"
)

var (
	listenAddress = flag.String("listen", "0.0.0.0:9090", "The address to listen on (host:port)")
)

type timeService struct {
	pb.UnimplementedTimeServiceServer
}

func initLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}

func (t *timeService) GetCurrentTime(ctx context.Context, req *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	slog.Info("Got time request")
	return &pb.GetCurrentTimeResponse{CurrentTime: time.Now().Format(time.RFC3339)}, nil
}

func main() {
	flag.Parse()
	initLogger()

	slog.Info("Time service starting", "address", *listenAddress)
	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}
	s := grpc.NewServer()

	ts := &timeService{}
	pb.RegisterTimeServiceServer(s, ts)

	if err := s.Serve(lis); err != nil {
		slog.Error("Failed to serve", "error", err)
		os.Exit(1)
	}
}
