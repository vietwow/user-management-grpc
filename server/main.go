package main

import (
    // "fmt"
    "time"
    "net"
    // "log"
    "go.uber.org/zap"

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

    "github.com/vietwow/user-management-grpc/pkg/logger"
    "github.com/vietwow/user-management-grpc/pkg/protocol/grpc/middleware"

    "github.com/vietwow/user-management-grpc/user-management-grpc/server/repository"

    uuid "github.com/satori/go.uuid"
)

type UserService struct {
    // db *pg.DB
    UserRepo repository.User
}

func NewUserService(db *pg.DB) *UserService {
    return &UserService{db: db}
}

func(s *UserService) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
    logger.Log.Info("Called function ListUsers()")
    var users []*pb.User
    // query := s.db.Model(&users).Order("id ASC")

    // err := query.Select()

    err := s.UserRepo.List(users)
    if err != nil {
        return nil, grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
    }

    return &pb.ListUserResponse{Users: users, Success: true}, nil
}

func(s *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    in.User.Id = uuid.NewV4().String()
    logger.Log.Info("Called function CreateUser() - Received:", zap.String("in.User.Id",in.User.Id)) 

    // err := s.db.Insert(in.User)
    err := s.UserRepo.Insert(in.User)
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
    logger.Log.Info("Called function CreateUsers() - Received:", zap.Strings("ids",ids))

    err := s.UserRepo.InsertBulk(&in.Users)
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not insert users into the database: %s", err)
    }

    return &pb.CreateUsersResponse{Ids: ids, Success: true}, nil
}

func(s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    logger.Log.Info("Called function GetUser() - Received:", zap.String("in.Id",in.Id))

    // var user pb.User
    // err := s.db.Model(&user).Where("id = ?", in.Id).First()
    user, err := s.UserRepo.Get(in.Id)
    if err != nil {
        return nil, grpc.Errorf(codes.NotFound, "Could not retrieve user from the database: %s", err)
    }

    return &pb.GetUserResponse{User: &user}, nil
}

func(s *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    logger.Log.Info("Called function UpdateUser() - Received:", zap.String("in.User.Id",in.User.Id))

    // res, err := s.db.Model(in.User).Column("username", "email", "password", "phone").WherePK().Update()

    err := s.UserRepo.Update(in.User)
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
    logger.Log.Info("Called function UpdateUsers() - Received:", zap.Strings("ids",ids))

    // res, err := s.db.Model(&in.Users).Column("username", "email", "password", "phone").WherePK().Update()
    err := s.GetUserResponse.UpdateBulk(&in.Users)
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Could not update users from the database: %s", err)
    }

    return &pb.UpdateUsersResponse{Ids: ids, Success: true}, nil
}

func(s *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    logger.Log.Info("Called function DeleteUser() - Received:", zap.String("in.Id",in.Id))

    // err := s.db.Delete(&pb.User{Id: in.Id})
    err := s.UserRepo.Delete(&pb.User{Id: in.Id})
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


    // initialize logger
    LogLevel := 0 // only print warn, not print info
    LogTimeFormat := "2006-01-02T15:04:05.999999999Z07:00"
    if err := logger.Init(LogLevel, LogTimeFormat); err != nil {
        logger.Log.Fatal("failed to initialize logger:", zap.String("reason", err.Error()))
    }

    //
    listen, err := net.Listen("tcp", port)
    if err != nil {
        logger.Log.Fatal("failed to listen:", zap.String("reason", err.Error()))
    }


    logger.Log.Info("Connecting PostgreSQL....")
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

    logger.Log.Info("Successfull Connected!")

    // Create Table from User struct generated by gRPC
    err = db.CreateTable(&pb.User{}, &orm.CreateTableOptions{
        IfNotExists:   true,
        FKConstraints: true,
    })
    if err != nil {
        logger.Log.Fatal("Create Table Failed:", zap.String("reason", err.Error()))
    }


    // gRPC server statup options
    opts := []grpc.ServerOption{}

    // add middleware
    opts = middleware.AddLogging(logger.Log, opts)

    // Creates a new gRPC server
    s := grpc.NewServer(opts...)
    
    userRep := repository.UserImpl{DB: db}

    // pb.RegisterUserServiceServer(s, &UserService{})
    pb.RegisterUserServiceServer(s, NewUserService(userRep))

    // graceful shutdown
    ctx := context.Background()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for range c {
            // sig is a ^C, handle it
            logger.Log.Warn("shutting down gRPC server...")

            s.GracefulStop()

            <-ctx.Done()
        }
    }()

    // start gRPC server
    logger.Log.Info("starting gRPC server...")
    if err := s.Serve(listen); err != nil {
        logger.Log.Fatal("failed to serve:", zap.String("reason", err.Error()))
    }
}
