package main

import (
    // "fmt"
    "time"
    "net"
    "log"
    "os"
    "os/signal"
    "io/ioutil"
    "bytes"

    "github.com/spf13/viper"

    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    // "google.golang.org/grpc/status"

    pb "github.com/vietwow/user-management-grpc/user"

    uuid "github.com/satori/go.uuid"
)

type UserService struct {
    db *pg.DB
}

func NewUserService(db *pg.DB) *UserService {
    return &UserService{db: db}
}

func(s *UserService) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
    var users []*pb.User
    query := s.db.Model(&users).Order("id ASC")

    err := query.Select()
    if err != nil {
        return nil, grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
    }

    return &pb.ListUserResponse{Users: users, Success: true}, nil
}

func(s *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    in.User.Id = uuid.NewV4().String()
    log.Printf("Received: %v", in.User.Id)

    // in.User.Id = uuid.NewV4().String()
    err := s.db.Insert(in.User)
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not insert user into the database: %s", err)
    }

    return &pb.CreateUserResponse{Id: in.User.Id, Success: true}, nil
}

func(s *UserService) CreateUsers(ctx context.Context, in *pb.CreateUsersRequest) (*pb.CreateUsersResponse, error) {
    var ids []string
    // fmt.Println(in.Users)
    for _, User := range in.Users {
        // fmt.Println(Users)

        User.Id = uuid.NewV4().String()
        // fmt.Println(User.Id)
        ids = append(ids, User.Id)
    }
    log.Printf("Received: %v", ids)

    err := s.db.Insert(&in.Users)
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not insert users into the database: %s", err)
    }

    return &pb.CreateUsersResponse{Ids: ids, Success: true}, nil
}

func(s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    log.Printf("Received: %v", in.Id)

    var user pb.User
    err := s.db.Model(&user).Where("id = ?", in.Id).First()
    if err != nil {
        return nil, grpc.Errorf(codes.NotFound, "Could not retrieve user from the database: %s", err)
    }

    return &pb.GetUserResponse{User: &user}, nil
}

func(s *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    log.Printf("Received: %v", in.User.Id)

    res, err := s.db.Model(in.User).Column("username", "email", "password", "phone").WherePK().Update()

    if res.RowsAffected() == 0 {
        return nil, grpc.Errorf(codes.NotFound, "Could not update user: not found")
    }
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not update user from the database: %s", err)
    }

    return &pb.UpdateUserResponse{Id: in.User.Id, Success: true}, nil
}

func(s *UserService) UpdateUsers(ctx context.Context, in *pb.UpdateUsersRequest) (*pb.UpdateUsersResponse, error) {
    var ids []string
    for _, User := range in.Users {
        ids = append(ids, User.Id)
    }
    log.Printf("Received: %v", ids)

    res, err := s.db.Model(&in.Users).Column("username", "email", "password", "phone").WherePK().Update()

    if res.RowsAffected() == 0 {
        return nil, grpc.Errorf(codes.NotFound, "Could not update users: not found")
    }
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not update users from the database: %s", err)
    }

    return &pb.UpdateUsersResponse{Ids: ids, Success: true}, nil
}

func(s *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    log.Printf("Received: %v", in.Id)

    err := s.db.Delete(&pb.User{Id: in.Id})
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not delete user from the database: %s", err)
    }

    return &pb.DeleteUserResponse{Id: in.Id, Success: true}, nil
}


// Config is configuration for Server
const (
    port = ":50051"
)

// default config in buffer

// toml
// var configDefault = []byte(`
// [database]
//     hostname = "localhost"
//     username = "root"
//     password = "khongbiet"
// `)

// yaml
// var yamlExample = []byte(`
// Hacker: true
// name: steve
// hobbies:
// - skateboarding
// - snowboarding
// - go
// clothing:
//      jacket: leather
//   trousers: denim
// age: 35
// eyes : brown
// beard: true
// `)

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
    viper.SetDefault("DatastoreDBUser", "postgres")
    viper.SetDefault("DatastoreDBPassword", "newhacker")
    viper.SetDefault("DatastoreDBHost", "localhost:5432")
    viper.SetDefault("DatastoreDBSchema", "grab")

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

func main () {
    // get configuration
    initConfig()

    DatastoreDBUser     := viper.GetString("DatastoreDBUser")
    DatastoreDBPassword := viper.GetString("DatastoreDBPassword")
    DatastoreDBHost     := viper.GetString("DatastoreDBHost")
    DatastoreDBSchema   := viper.GetString("DatastoreDBSchema")


    //
    listen, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }


    log.Println("Connecting PostgreSQL....")
    // Connect to PostgresQL
    db := pg.Connect(&pg.Options{
        User:     DatastoreDBUser,
        Password: DatastoreDBPassword,
        Database: DatastoreDBSchema,
        Addr:     DatastoreDBHost,
        RetryStatementTimeout: true,
        MaxRetries:            4,
        MinRetryBackoff:       250 * time.Millisecond,
    })

    defer db.Close()

    log.Println("Successfull Connected!")

    // Create Table from User struct generated by gRPC
    err = db.CreateTable(&pb.User{}, &orm.CreateTableOptions{
        IfNotExists:   true,
        FKConstraints: true,
    })
    if err != nil {
        log.Fatalf("Create Table Failed: %v",err)
    }

    // Creates a new gRPC server
    s := grpc.NewServer()
    // pb.RegisterUserServiceServer(s, &UserService{})
    pb.RegisterUserServiceServer(s, NewUserService(db))

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
