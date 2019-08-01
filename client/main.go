package main

import (
    "context"
    "log"
    "time"

    "gopkg.in/alecthomas/kingpin.v2"

    "google.golang.org/grpc"

    pb "github.com/vietwow/user-management-grpc/user"
)

var (
    address  = kingpin.Flag("server", "gRPC server in format host:port").Envar("SERVER").String()
    version  = "0.0.0"

    verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
)

func main() {
    // get configuration
    kingpin.Version(version)
    kingpin.Parse()

    // Set up a connection to the server.
    conn, err := grpc.Dial(*address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewUserServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()


    // Call Create
    username := "vietwow"
    email    := "vietwow@gmail.com"
    password := "newhacker"
    phone    := "123456"

    req1 := pb.CreateUserRequest{
        User: &pb.User{
            Username: "Username (" + username + ")",
            Email:    "Email (" + email + ")",
            Password: "Password (" + password + ")",
            Phone:    "Phone (" + phone + ")",
        },
    }
    res1, err := c.CreateUser(ctx, &req1)
    if err != nil {
        log.Fatalf("CreateUser failed: %v", err)
    }
    log.Printf("CreateUser result: <%+v>\n\n", res1)

    // Call GetUser
    id := res1.UserId

    req2 := pb.GetUserRequest{
        UserId: id,
    }
    res2, err := c.GetUser(ctx, &req2)
    if err != nil {
        log.Fatalf("GetUser failed: %v", err)
    }
    log.Printf("GetUser result: <%+v>\n\n", res2)

    // Call UpdateUser
    req3 := pb.UpdateUserRequest{
        User: &pb.User{
            UserId:   res2.User.UserId,
            Username: res2.User.Username,
            Email:    res2.User.Email,
            Password: res2.User.Password,
            Phone:    res2.User.Phone,
        },
    }
    res3, err := c.UpdateUser(ctx, &req3)
    if err != nil {
        log.Fatalf("UpdateUser failed: %v", err)
    }
    log.Printf("UpdateUser result: <%+v>\n\n", res3)

    // Call ListUser
    req4 := pb.ListUserRequest{}
    res4, err := c.ListUser(ctx, &req4)
    if err != nil {
        log.Fatalf("ListUser failed: %v", err)
    }
    log.Printf("ListUser result: <%+v>\n\n", res4)

    // Call DeleteUser
    req5 := pb.DeleteUserRequest{
        UserId:  id,
    }
    res5, err := c.DeleteUser(ctx, &req5)
    if err != nil {
        log.Fatalf("DeleteUser failed: %v", err)
    }
    log.Printf("DeleteUser result: <%+v>\n\n", res5)
}