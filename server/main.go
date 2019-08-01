package main

import (
    "fmt"
    "net"
    "log"
    "os"
    "os/signal"

    "gopkg.in/alecthomas/kingpin.v2"

    "database/sql"

    // mysql driver
    _ "github.com/go-sql-driver/mysql"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    pb "github.com/vietwow/user-management-grpc/user"
)

type UserService struct {
    db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{db: db}
}

// connect returns SQL database connection from the pool
func (s *UserService) connect(ctx context.Context) (*sql.Conn, error) {
    c, err := s.db.Conn(ctx)
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
    }
    return c, nil
}

func(s *UserService) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
    // get SQL connection from pool
    c, err := s.connect(ctx)
    if err != nil {
        return nil, err
    }
    defer c.Close()

    // get User list
    rows, err := c.QueryContext(ctx, "SELECT `UserId`, `Username`, `Email`, `Password`, `Phone` FROM User")
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to select from User-> "+err.Error())
    }
    defer rows.Close()

    var list []*pb.User
    for rows.Next() {
        user := new(pb.User)
        if err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Phone); err != nil {
            return nil, status.Error(codes.Unknown, "failed to retrieve field values from User row-> "+err.Error())
        }
        list = append(list, user)
    }

    if err := rows.Err(); err != nil {
        return nil, status.Error(codes.Unknown, "failed to retrieve data from User-> "+err.Error())
    }

    return &pb.ListUserResponse{Users: list}, nil
}

func(s *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    log.Printf("Received: %v", in.User.UserId)

    // get SQL connection from pool
    c, err := s.connect(ctx)
    if err != nil {
        return nil, err
    }
    defer c.Close()

    // insert ToDo entity data
    res, err := c.ExecContext(ctx, "INSERT INTO User(`UserId`, `Username`, `Email`, `Password`, `Phone`) VALUES(?, ?, ?, ?, ?)",
        in.User.UserId, in.User.Username, in.User.Email, in.User.Password, in.User.Phone)
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to insert into User-> "+err.Error())
    }

    // get ID of creates ToDo
    id, err := res.LastInsertId()
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to retrieve id for created User-> "+err.Error())
    }

    return &pb.CreateUserResponse{UserId: id, Success: true}, nil
}

func(s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    log.Printf("Received: %v", in.UserId)

    // get SQL connection from pool
    c, err := s.connect(ctx)
    if err != nil {
        return nil, err
    }
    defer c.Close()

    // query ToDo by ID
    rows, err := c.QueryContext(ctx, "SELECT `UserId`, `Username`, `Email`, `Password`, `Phone` FROM User WHERE `ID`=?",
        in.UserId)
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to select from User-> "+err.Error())
    }
    defer rows.Close()

    if !rows.Next() {
        if err := rows.Err(); err != nil {
            return nil, status.Error(codes.Unknown, "failed to retrieve data from User-> "+err.Error())
        }
        return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
            in.UserId))
    }

    // get ToDo data
    var user pb.User
    if err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Phone); err != nil {
        return nil, status.Error(codes.Unknown, "failed to retrieve field values from User row-> "+err.Error())
    }

    if rows.Next() {
        return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple User rows with ID='%d'",
            in.UserId))
    }

    return &pb.GetUserResponse{User: &user}, nil
}

func(s *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    log.Printf("Received: %v", in.User.UserId)

    // get SQL connection from pool
    c, err := s.connect(ctx)
    if err != nil {
        return nil, err
    }
    defer c.Close()

    // update User
    res, err := c.ExecContext(ctx, "UPDATE User SET `Email`=?, `Password`=?, `Phone`=? WHERE `ID`=?",
        in.User.Email, in.User.Password, in.User.Phone, in.User.UserId)
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to update User-> "+err.Error())
    }

    rows, err := res.RowsAffected()
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
    }

    if rows == 0 {
        return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
            in.User.UserId))
    }

    return &pb.UpdateUserResponse{UserId: in.User.UserId, Success: true}, nil
}

func(s *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    log.Printf("Received: %v", in.UserId)

    // get SQL connection from pool
    c, err := s.connect(ctx)
    if err != nil {
        return nil, err
    }
    defer c.Close()

    // delete User
    res, err := c.ExecContext(ctx, "DELETE FROM User WHERE `ID`=?", in.UserId)
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to delete User-> "+err.Error())
    }

    rows, err := res.RowsAffected()
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
    }

    if rows == 0 {
        return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
            in.UserId))
    }

    return &pb.DeleteUserResponse{UserId: in.UserId, Success: true}, nil
}


// Config is configuration for Server
const (
    port = ":50051"
)

var (
    DatastoreDBHost     = kingpin.Flag("db-host", "Database host").Envar("DB_HOST").String()
    DatastoreDBUser     = kingpin.Flag("db-user", "Database user").Envar("DB_USER").Default("root").String()
    DatastoreDBPassword = kingpin.Flag("db-password", "Database password").Envar("DB_PASSWORD").Default("newhacker").String()
    DatastoreDBSchema   = kingpin.Flag("db-schema", "Database schema").Envar("DB_SCHEMA").Default("").String()
    version             = "0.0.0"

    verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
)

func main() {
    // get configuration
    kingpin.Version(version)
    kingpin.Parse()


    //
    listen, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }


	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        *DatastoreDBUser,
        *DatastoreDBPassword,
        *DatastoreDBHost,
        *DatastoreDBSchema,
        param)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("failed to open database: %v", err)
    }
    defer db.Close()


    //
    user_service := NewUserService(db)

    // Creates a new gRPC server
    s := grpc.NewServer()
    // pb.RegisterUserServiceServer(s, &UserService{})
    pb.RegisterUserServiceServer(s, user_service)

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