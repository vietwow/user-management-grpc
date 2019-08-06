package main

import (
    // "fmt"
    "context"
    "log"
    "time"
    "io/ioutil"
    "bytes"

    "github.com/spf13/viper"

    "google.golang.org/grpc"

    pb "github.com/vietwow/user-management-grpc/user"
)

func initConfig() error {
    // log.Println("Loading config...")
    viper.SetConfigType("yaml")
    // viper.SetDefault("proxyList", "/etc/proxy.list")
    // viper.SetDefault("check", map[string]interface{}{
    //     "url":      "http://ya.ru",
    //     "string":   "yandex",
    //     "interval": "60m",
    //     "timeout":  "5s",
    // })
    viper.SetDefault("SERVER", "localhost:50051")

    configFile := "config.yaml"

    file, err := ioutil.ReadFile(configFile)
    if err != nil {
        return err
    }

    err = viper.ReadConfig(bytes.NewReader(file))
    if err != nil {
        return err
    }

    return nil
}

func main() {
    // get configuration
    initConfig()
    address := viper.GetString("SERVER")

    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewUserServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()


    // Call CreateUser
    username := "vietwow"
    email    := "vietwow@gmail.com"
    password := "newhacker"
    phone    := "123456"

    req1 := pb.CreateUserRequest{
        User: &pb.User{
            Username: username,
            Email:    email,
            Password: password,
            Phone:    phone,
        },
    }
    res1, err := c.CreateUser(ctx, &req1)
    if err != nil {
        log.Fatalf("CreateUser failed: %v", err)
    }
    log.Printf("CreateUser result: <%+v>\n\n", res1)

    // Call GetUser
    id := res1.Id

    req2 := pb.GetUserRequest{
        Id: id,
    }
    res2, err := c.GetUser(ctx, &req2)
    if err != nil {
        log.Fatalf("GetUser failed: %v", err)
    }
    log.Printf("GetUser result: <%+v>\n\n", res2)

    // Call UpdateUser
    req3 := pb.UpdateUserRequest{
        User: &pb.User{
            Id:   res2.User.Id,
            Username: res2.User.Username,
            Email:    res2.User.Email + " + updated",
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
        Id:  id,
    }
    res5, err := c.DeleteUser(ctx, &req5)
    if err != nil {
        log.Fatalf("DeleteUser failed: %v", err)
    }
    log.Printf("DeleteUser result: <%+v>\n\n", res5)

    // Call CreateUsers
    users := []*pb.User{
        {
            Username: "vietwow",
            Email:    "vietwow@gmail.com",
            Password: "newhacker",
            Phone:    "123456",
        },
        {
            Username: "vietwow2",
            Email:    "vietwow2@gmail.com",
            Password: "newhacker",
            Phone:    "123456",
        },
        {
            Username: "vietwow3",
            Email:    "vietwow3@gmail.com",
            Password: "newhacker",
            Phone:    "123456",
        },
    }

    req6 := pb.CreateUsersRequest{
        Users: users,
    }
    res6, err := c.CreateUsers(ctx, &req6)
    if err != nil {
        log.Fatalf("CreateUsers failed: %v", err)
    }
    log.Printf("CreateUsers result: <%+v>\n\n", res6)

    // List the users
    rlist, err := c.ListUser(
        context.Background(),
        &pb.ListUserRequest{},
    )

    // Call UpdateUsers
    req7 := pb.UpdateUsersRequest{
        Users: rlist.Users,
    }
    res7, err := c.UpdateUsers(ctx, &req7)
    if err != nil {
        log.Fatalf("UpdateUsers failed: %v", err)
    }
    log.Printf("UpdateUsers result: <%+v>\n\n", res7)
}