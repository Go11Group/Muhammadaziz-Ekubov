package postgres

import (
	"add/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

type UserFilter struct {
	Id       string
	Name     string
	Email    string
	Password string
	Limit    int
}

func (u *UserRepo) CreateUser(user *model.User) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		return errors.New("failed to parse birthday")
	}

	user.UserID = uuid.NewString()
	_, err = u.DB.Exec("insert into Users (user_id,name,email,birthday,password) VALUES ($1,$2,$3,$4,$5)",
		user.UserID, user.Name, user.Email, birthday, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) ReadUserByID(UserID string) (*model.User, error) {
	var user model.User
	row := u.DB.QueryRow("select * from Users where user_id = $1", UserID)

	err := row.Scan(
		&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) UpdateUser(User *model.User) error {
	birthday, err := time.Parse("2006-01-02", User.Birthday)
	if err != nil {
		return errors.New("failed to parse birthday")
	}
	_, err = u.DB.Exec("update Users set name = $1, email = $2, birthday = $3, password = $4 where user_id = $5",
		User.Name, User.Email, birthday, User.Password, User.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) DeleteUser(UserID string) error {
	now := time.Now()
	d := now.Unix()

	_, err := u.DB.Exec("update users set deleted_at=$1 where user_id = $2", d, UserID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetAll() ([]model.Users, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	filter := ""
	f := UserFilter{}
	query := `select user_id, , name, email, birthday,password,
	 	from Users where deleted_at = 0 `

	if len(f.Id) > 0 {
		params["id"] = f.Id
		filter += "and user_id = :user_id"
	}
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += "and name = :name "
	}
	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += "and email = :email "
	}
	if len(f.Password) > 0 {
		params["password"] = f.Password
		filter += "and password = :password"
	}

	query = query + filter + limit

	query, arr = ReplaceQueryParams(query, params)
	db := &sql.DB{}
	rows, err := db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		user := model.Users{}

		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
