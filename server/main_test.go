package main

import (
    "context"
    "testing"

    "github.com/spf13/viper"
    "time"

    api "github.com/vietwow/user-management-grpc/user"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type UserSuite struct {
    suite.Suite
    User *UserService
}

func TestUserTestSuite(t *testing.T) {
    // get configuration
    initConfig()

    DatastoreDBUser     := viper.GetString("DatastoreDBUser")
    DatastoreDBPassword := viper.GetString("DatastoreDBPassword")
    DatastoreDBHost     := viper.GetString("DatastoreDBHost")
    DatastoreDBSchema   := viper.GetString("DatastoreDBSchema")

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

    suite.Run(t, &UserSuite{
        User: &UserService{db: db},
    })
}

// It will be run before each test of the suite. Set the default value etc. here.
func (s *UserSuite) SetupTest() {
    s.User.db.DropTable(&api.User{}, &orm.DropTableOptions{IfExists: true})
    s.User.db.CreateTable(&api.User{}, nil)
}

// Run after each test in the suite.
func (s *UserSuite) TearDownTest() {
    s.User.db.DropTable(&api.User{}, &orm.DropTableOptions{IfExists: true})
}

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
    assert.NotEqual(s.T(), rcreate.Id, "")
}

func (s *UserSuite) TestCreateUsers() {
    rcreate, err := s.User.CreateUsers(
        context.Background(),
        &api.CreateUsersRequest{
            Users: []*api.User{
                &api.User{
                    Username: "vietwow",
                    Email:    "vietwow@gmail.com",
                    Password: "newhacker",
                    Phone:    "123456",
                },
                &api.User{
                    Username: "vietwow2",
                    Email:    "vietwow2@gmail.com",
                    Password: "newhacker",
                    Phone:    "123456",
                },
            },
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rcreate)
    for _, id := range rcreate.Ids {
        assert.NotEqual(s.T(), id, "")
    }
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
    assert.NotEqual(s.T(), rcreate.Id, "")

    id := rcreate.Id

    rget, err := s.User.GetUser(
        context.Background(),
        &api.GetUserRequest{
            Id: id,
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
    assert.NotEqual(s.T(), rcreate.Id, "")

    id := rcreate.Id

    rdel, err := s.User.DeleteUser(
        context.Background(),
        &api.DeleteUserRequest{
            Id: id,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rdel)

    // Getting the user should fail this time
    rget, err := s.User.GetUser(
        context.Background(),
        &api.GetUserRequest{
            Id: id,
        },
    )
    assert.Nil(s.T(), rget)
    assert.NotNil(s.T(), err)
    assert.Contains(s.T(), err.Error(), "Could not retrieve user from the database: pg: no rows in result set")
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
    assert.NotEqual(s.T(), rcreate.Id, "")

    id := rcreate.Id

    newUser := &api.User{
        Id: id,
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

    // Getting the user should return the updated version
    rget, err := s.User.GetUser(
        context.Background(),
        &api.GetUserRequest{
            Id: id,
        },
    )
    assert.NotNil(s.T(), rget)
    assert.Nil(s.T(), err)
    assert.Equal(s.T(), rget.User.Username, newUser.Username)
    assert.Equal(s.T(), rget.User.Email, newUser.Email)
    assert.Equal(s.T(), rget.User.Password, newUser.Password)
    assert.Equal(s.T(), rget.User.Phone, newUser.Phone)
}

func (s *UserSuite) TestUpdateUsers() {
    users := []*api.User{
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

    // Create the users
    resp, err := s.User.CreateUsers(
        context.Background(),
        &api.CreateUsersRequest{
            Users: users,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), resp)

    // List the Users and update their fields
    rlist, err := s.User.ListUser(
        context.Background(),
        &api.ListUserRequest{},
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rlist)
    assert.NotNil(s.T(), rlist.Users)

    rupdate, err := s.User.UpdateUsers(
        context.Background(),
        &api.UpdateUsersRequest{
            Users: rlist.Users,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rupdate)

    // List again and see if the entries have had their fields changed
    rlist, err = s.User.ListUser(
        context.Background(),
        &api.ListUserRequest{},
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rlist)
    assert.NotNil(s.T(), rlist.Users)
}

func (s *UserSuite) TestListUser() {
    users := []*api.User{
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

    // List with empty database
    rlist, err := s.User.ListUser(
        context.Background(),
        &api.ListUserRequest{},
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rlist)
    assert.Nil(s.T(), rlist.Users)
    assert.Equal(s.T(), len(rlist.Users), 0)

    // Create the User Users
    rcreate, err := s.User.CreateUsers(
        context.Background(),
        &api.CreateUsersRequest{
            Users: users,
        },
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rcreate)

    // List the Users
    rlist, err = s.User.ListUser(
        context.Background(),
        &api.ListUserRequest{},
    )
    assert.Nil(s.T(), err)
    assert.NotNil(s.T(), rlist)
    assert.NotNil(s.T(), rlist.Users)
    assert.Equal(s.T(), len(rlist.Users), 3)
}