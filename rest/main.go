/*

curl http://localhost:8080/api/v1/users

curl -X POST -k http://localhost:8080/api/v1/users -d '{"username": " huyen", "email": "huyen@gmail.com", "password": "123123", "phone":"0905023639"}'
{"id":"07064352-560a-4885-8919-db566c1b0372","success":true}%

*/

package main

import (
    // "flag"
    "net/http"
    "os"
    "os/signal"
    // "log"
    "go.uber.org/zap"

    pb "github.com/vietwow/user-management-grpc/user"

    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "golang.org/x/net/context"
    "google.golang.org/grpc"

    "github.com/vietwow/user-management-grpc/pkg/logger"
    "github.com/vietwow/user-management-grpc/pkg/protocol/rest/middleware"
)

var (
    grpcPort = ":50051"
    httpPort = ":8080"
)


func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
    if err != nil {
        logger.Log.Fatal("failed to start HTTP gateway:", zap.String("reason", err.Error()))
    }

    srv := http.Server{
        Addr:    httpPort,
        // Handler: mux,
        Handler: middleware.AddRequestID(middleware.AddLogger(logger.Log, mux)),
    }

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for range c {
            logger.Log.Warn("shutting down HTTP server...")
            if err := srv.Shutdown(context.Background()); err != nil {
                logger.Log.Fatal("failed to shutdown HTTP server:", zap.String("reason", err.Error()))
            }
        }
    }()

    srv.ListenAndServe()
}
