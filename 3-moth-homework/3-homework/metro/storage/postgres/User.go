package postgres

import (
	"Muhammadaziz-Ekubov/3-moth-homework/3-homework/metro/model"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetAll() ([]*model.Users, error) {
	rows, err := u.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	var users []*model.Users
	for rows.Next() {
		var user model.Users

		if user.DeletedAt == 0 {
			err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
			if err != nil {
				return nil, err
			}
			users = append(users, &user)
		}
	}

	return users, nil
}

func (u *UserRepository) GetByID(id string) (*model.Users, error) {
	var user model.Users

	err := u.db.QueryRow("select id, name, phone, age from users where id =$1", id).
		Scan(&user.ID, &user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) CreateUser(user *model.Users) error {

	user.ID = uuid.NewString()

	_, err := u.db.Exec("insert into users(id, name, phone, age) values ($1, $2, $3)", user.ID, user.Name, user.Phone, user.Age)
	if err != nil {
		return err
	}
	return nil

}

func (u *UserRepository) UpdateUser(user *model.Users) error {
	_, err := u.db.Exec("update users set name=$1 , phone=$2, age = $3 where id=$4", user.Name, user.Phone, user.Age, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteUser(id string) error {

	now := time.Now()
	DeletedAt := now.Unix()

	_, err := u.db.Exec("update from users set deleted_at=$1 where id=$2", DeletedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetBalanceByID(id string) (*model.Users, error) {
	var user model.Users

	_, err := u.db.Query("select u.id, u.name, u.phone, u.age c.amount from users  u, card c where id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
