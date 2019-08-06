package main

import (
    "context"
    "testing"

    "database/sql"

    "github.com/spf13/viper"
    "fmt"

    api "github.com/vietwow/user-management-grpc/user"
    // "github.com/go-pg/pg"
    // "github.com/go-pg/pg/orm"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type UserSuite struct {
    suite.Suite
    User *UserService
}

var (
	ctx context.Context
	db  *sql.DB
)

func TestUserTestSuite(t *testing.T) {
    // db := pg.Connect(&pg.Options{
    //     User:     "postgres",
    //     Database: "User",
    //     Addr:     "localhost:5432",
    //     RetryStatementTimeout: true,
    //     MaxRetries:            4,
    //     MinRetryBackoff:       250 * time.Millisecond,
    // })

    suite.Run(t, &UserSuite{
        User: &UserService{db: db},
    })
}

// It will be run before each test of the suite. Set the default value etc. here.
func (s *UserSuite) SetupTest() {
    // get configuration
    initConfig()

    DatastoreDBUser     := viper.GetString("DatastoreDBUser")
    DatastoreDBPassword := viper.GetString("DatastoreDBPassword")
    DatastoreDBHost     := viper.GetString("DatastoreDBHost")
    DatastoreDBSchema   := viper.GetString("DatastoreDBSchema")

    param := "parseTime=true"

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        DatastoreDBUser,
        DatastoreDBPassword,
        DatastoreDBHost,
        DatastoreDBSchema,
        param)
    db, _ := sql.Open("mysql", dsn)
    defer db.Close()
    // get SQL connection from pool
    c, _ := db.Conn(ctx)

    //res, err := c.ExecContext(ctx,"DROP TABLE User")
    //res, err := c.ExecContext(ctx,"CREATE TABLE User (UserId BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT, Username VARCHAR(200) DEFAULT NUL, Email VARCHAR(1024) DEFAULT NULL, Password VARCHAR(1024) DEFAULT NULL, Phone VARCHAR(1024) DEFAULT NULL")
    c.ExecContext(ctx,"Truncate User")
}

// Run after each test in the suite.
func (s *UserSuite) TearDownTest() {}

func (s *UserSuite) TestCreateUser() {
    rcreate, err := s.User.CreateUser(
        context.Background(),
        &api.CreateUserRequest{
            User: &api.User{
		        Username: "vietwow",
		        Email:    "vietwow@gmail.com",
		        Password: "newhacker",
		        Phone:    "123456",
            },
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rcreate)
    assert.NotEqual(s.T(), rcreate.UserId, "")
}

func (s *UserSuite) TestGetUser() {
    user := &api.User{
        Username: "vietwow",
        Email:    "vietwow@gmail.com",
        Password: "newhacker",
        Phone:    "123456",
    }

    rcreate, err := s.User.CreateUser(
        context.Background(),
        &api.CreateUserRequest{
            User: user,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rcreate)
    assert.NotEqual(s.T(), rcreate.UserId, "")

    id := rcreate.UserId

    rget, err := s.User.GetUser(
        context.Background(),
        &api.GetUserRequest{
            UserId: id,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rget)
    assert.NotNil(s.T(), rget.User)
    assert.Equal(s.T(), rget.User, user)
}

func (s *UserSuite) TestDeleteUser() {
    user := &api.User{
        Username: "vietwow",
        Email:    "vietwow@gmail.com",
        Password: "newhacker",
        Phone:    "123456",
    }

    rcreate, err := s.User.CreateUser(
        context.Background(),
        &api.CreateUserRequest{
            User: user,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rcreate)
    assert.NotEqual(s.T(), rcreate.UserId, "")

    id := rcreate.UserId

    rdel, err := s.User.DeleteUser(
        context.Background(),
        &api.DeleteUserRequest{
            UserId: id,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rdel)

    // Getting the User item should fail this time
    rget, err := s.User.GetUser(
        context.Background(),
        &api.GetUserRequest{
            UserId: id,
        },
    )
    assert.Nil(s.T(), rget)
    assert.NotNil(s.T(), err)
    assert.Contains(s.T(), err.Error(), "Could not retrieve item from the database: pg: no rows in result set")
}

func (s *UserSuite) TestUpdateUser() {
    user := &api.User{
        Username: "vietwow",
        Email:    "vietwow@gmail.com",
        Password: "newhacker",
        Phone:    "123456",
    }

    rcreate, err := s.User.CreateUser(
        context.Background(),
        &api.CreateUserRequest{
            User: user,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rcreate)
    assert.NotEqual(s.T(), rcreate.UserId, "")

    id := rcreate.UserId

    newUser := &api.User{
        Username: "vietwow2",
        Email:    "vietwow2@gmail.com",
        Password: "newhacker",
        Phone:    "123456",
    }

    rupdate, err := s.User.UpdateUser(
        context.Background(),
        &api.UpdateUserRequest{
            User: newUser,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rupdate)

    // Getting the User item should return the updated version
    rget, err := s.User.GetUser(
        context.Background(),
        &api.GetUserRequest{
            UserId: id,
        },
    )
    assert.NotNil(s.T(), rget)
    assert.Nil(s.T(), err)
    assert.Equal(s.T(), rget.User.Username, newUser.Username)
    assert.Equal(s.T(), rget.User.Email, newUser.Email)
    assert.Equal(s.T(), rget.User.Password, newUser.Password)
    assert.Equal(s.T(), rget.User.Phone, newUser.Phone)
}


	// func (s *UserSuite) TestListUser() {
	//     items := []*api.User{
	//         {
	//             Title:       "item_1",
	//             Description: "item desc 1",
	//             Completed:   true,
	//         },
	//         {
	//             Title:       "item_2",
	//             Description: "item desc 2",
	//         },
	//         {
	//             Title:       "item_3",
	//             Description: "item desc 3",
	//         },
	//         {
	//             Title:       "item_4",
	//             Description: "item desc 4",
	//             Completed:   true,
	//         },
	//     }

	//     // List with empty database
	//     rlist, err := s.User.ListUser(
	//         context.Background(),
	//         &api.ListUserRequest{},
	//     )
	//     assert.Nil(s.T(), err)
	//     assert.NotNil(s.T(), rlist)
	//     assert.Nil(s.T(), rlist.Items)
	//     assert.Equal(s.T(), len(rlist.Items), 0)

	//     // Create the User items
	//     rcreate, err := s.User.CreateUsers(
	//         context.Background(),
	//         &api.CreateUsersRequest{
	//             Items: items,
	//         },
	//     )
	//     assert.Nil(s.T(), err)
	//     assert.NotNil(s.T(), rcreate)

	//     // List the items
	//     rlist, err = s.User.ListUser(
	//         context.Background(),
	//         &api.ListUserRequest{},
	//     )
	//     assert.Nil(s.T(), err)
	//     assert.NotNil(s.T(), rlist)
	//     assert.NotNil(s.T(), rlist.Items)
	//     assert.Equal(s.T(), len(rlist.Items), 4)

	//     // Limit the result of List
	//     rlist, err = s.User.ListUser(
	//         context.Background(),
	//         &api.ListUserRequest{
	//             Limit: 2,
	//         },
	//     )
	//     assert.Nil(s.T(), err)
	//     assert.NotNil(s.T(), rlist)
	//     assert.NotNil(s.T(), rlist.Items)
	//     assert.Equal(s.T(), len(rlist.Items), 2)

	//     // Only list non completed items
	//     rlist, err = s.User.ListUser(
	//         context.Background(),
	//         &api.ListUserRequest{
	//             NotCompleted: true,
	//         },
	//     )
	//     assert.Nil(s.T(), err)
	//     assert.NotNil(s.T(), rlist)
	//     assert.NotNil(s.T(), rlist.Items)
	//     assert.Equal(s.T(), len(rlist.Items), 2)
	// }