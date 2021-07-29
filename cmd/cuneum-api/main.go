package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/cuneum/cuneum-api/internal/handler"
	"github.com/cuneum/cuneum-api/internal/service"
	"github.com/cuneum/cuneum-api/pkg/pb"
)

const (
	httpAddress = ":8080"
	grpcAddress = ":8081"
)

func main() {
	grpcS := grpcServer()
	httpS := httpServer()
	log.Println("start")

	ctx, f := signal.NotifyContext(context.Background(), os.Interrupt)
	defer f()

	<-ctx.Done()
	httpS.Shutdown(context.Background())
	grpcS.GracefulStop()
	log.Println("shutdown")
}

func grpcServer() *grpc.Server {
	l, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, handler.NewArticle(service.NewArticleService()))
	health.RegisterHealthServer(server, handler.NewHealth())
	go func() {
		if err := server.Serve(l); err != nil {
			panic(err)
		}
	}()

	return server
}

func httpServer() *http.Server {
	mux := runtime.NewServeMux()

	err := pb.RegisterArticleServiceHandlerFromEndpoint(context.Background(), mux, grpcAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    httpAddress,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return server
}
