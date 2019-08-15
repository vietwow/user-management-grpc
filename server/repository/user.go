package repository

import (
	"github.com/go-pg/pg"
	"github.com/xuanit/testing/todo/pb"
)

type UserImpl struct {
	DB *pg.DB
}

type User interface {
	List() ([]*pb.User, error)
	Get(id string) (pb.User, error)
	Insert(user pb.User) (error)
	InsertBulk(users []*pb.User) (error)
	Update(user pb.User) (error)
	UpdateBulk(users []*pb.User) (error)
	Delete(id string) (error)
}

// func (r UserImpl) List(limit int32, notCompleted bool) ([]*pb.Todo, error) {
// 	var items []*pb.Todo
// 	query := r.DB.Model(&items).Order("created_at ASC")
// 	if limit > 0 {
// 		query.Limit(int(limit))
// 	}
// 	if notCompleted {
// 		query.Where("completed = false")
// 	}
// 	err := query.Select()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// func (r UserImpl) Insert(items *pb.Todo)  error {
// 	err := r.DB.Insert(items)
// 	return err
// }

// func (r UserImpl) Get(id string) (*pb.Todo, error) {
// 		var item pb.Todo
// 		err := r.DB.Model(&item).Where("id = ?", id).First()
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &item, nil
// }

// func (r UserImpl) Delete(id string) error {
// 	err := r.DB.Delete(&pb.Todo{Id: id})
// 	return err
// }

func (r UserImpl) List() ([]*pb.User, error) {
	var users []*pb.User
	query := r.DB.Model(&users).Order("id ASC")
	// if limit > 0 {
	// 	query.Limit(int(limit))
	// }
	// if notCompleted {
	// 	query.Where("completed = false")
	// }
	err := query.Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserImpl) Get(id string) (*pb.User, error) {
		var user pb.User
		err := r.DB.Model(&user).Where("id = ?", id).First()
		if err != nil {
			return nil, err
		}
		return &user, nil
}


func (r UserImpl) Insert(user *pb.User) error {
	err := r.DB.Insert(user)

    if err != nil {
        return nil, log.Fatal("Could not insert user into the database: %s", err)
    }
	return err
}

func (r UserImpl) InsertBulk(users []*pb.User) error {
	err := r.DB.Insert(users)

    if err != nil {
        return nil, log.Fatal("Could not insert users into the database: %s", err)
    }
	return err
}


func (r UserImpl) Update(user *pb.User) error {
	res, err := r.DB.Model(user).Column("username", "email", "password", "phone").WherePK().Update()

    if res.RowsAffected() == 0 {
        return nil, log.Fatal("Could not update user: not found")
    }
    if err != nil {
        return nil, log.Fatal("Could not update user from the database: %s", err)
    }
	return err
}


func (r UserImpl) UpdateBulk(users []*pb.User) error {
	res, err := r.DB.Model(&users).Column("username", "email", "password", "phone").WherePK().Update()

    if res.RowsAffected() == 0 {
        return nil, log.Fatal("Could not update users: not found")
    }
    if err != nil {
        return nil, log.Fatal("Could not update users from the database: %s", err)
    }
	return err
}


func (r UserImpl) Delete(id string) error {
	err := r.DB.Delete(&pb.User{Id: id})
	return err
}
