/*

curl http://localhost:8080/api/v1/users

curl -X POST -k http://localhost:8080/api/v1/users -d '{"name": " world"}

*/

package main

import (
    "flag"
    "net/http"
    "os"
    "os/signal"
    "log"

    pb "github.com/vietwow/user-management-grpc/user"

    "github.com/golang/glog"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

var (
    grpcPort = ":50051"
    httpPort = ":8080"
)


func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
    if err != nil {
        log.Fatalf("failed to start HTTP gateway %v", err)
    }

    srv := http.Server{
        Addr:    httpPort,
        Handler: mux,
    }

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for range c {
            log.Println("shutting down HTTP server...")
            if err := srv.Shutdown(context.Background()); err != nil {
                log.Fatalf("failed to shutdown HTTP server: %v", err)
            }
        }
    }()

    return srv.ListenAndServe()
}

func main() {
    flag.Parse()
    defer glog.Flush()

    if err := run(); err != nil {
        glog.Fatal(err)
    }
}
