package main

import (
    // "fmt"
    "net"
    "log"
	"os"
	"os/signal"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    
    pb "github.com/vietwow/user-management-grpc/user"
)

const (
    port = ":50051"
)

type UserService struct {}

// func NewUserService() UserService {
// 	return &UserService{}	
// }

// func(s *server) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
//     // for _, user := range s.savedUsers {
//     //     ...
//     // }
//     userId := 1
//     return &pb.ListUserResponse{Id: userId, Success: true}, nil
// }

func(s *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    log.Printf("Received: %v", in.User.UserId)

	return &pb.CreateUserResponse{UserId: in.User.UserId, Success: true}, nil
}

// func(s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
//     log.Printf("Received: %v", in.Username)

// 	return &pb.GetUserResponse{User: {UserId: 1, Username: 'vietwow', Email: 'vietwow@gmail.com', Password: '123456', Pshone: '123456'}}, nil
// }

func(s *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    log.Printf("Received: %v", in.Username)

	return &pb.DeleteUserResponse{UserId: 123, Success: true}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	// service := pb.UserServiceServer(&UserService{})
	pb.RegisterUserServiceServer(s, &UserService{})

	// graceful shutdown
	ctx := context.Background()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			s.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	if err := s.Serve(listen); err != nil {
	    log.Fatalf("failed to serve: %v", err)
	}
}