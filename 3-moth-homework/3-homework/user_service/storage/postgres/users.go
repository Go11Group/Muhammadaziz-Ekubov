package postgres

import (
	"atto/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserRepo struct {
	DB *sql.DB
}

// create
func (u *UserRepo) CreateUser(user *models.CreateUser) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into users(id, name, phone, age, created_at) 
	values ($1,$2,$3,$4,$5)`
	_, err = tx.Exec(query, uuid.NewString(), user.Name, user.Phone, user.Age,
		time.Now())
	if err != nil {
		return err
	}
	return nil
}

// read all
func (u *UserRepo) GetAllUser(filter models.UserFilter) (*[]models.User, error) {
	query := "select * from users where and deleted_at is null"
	params := []interface{}{}
	paramcount := 1

	if filter.Name != nil {
		query += " and name = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.Name)
		paramcount++
	}
	if filter.Phone != nil {
		query += " and phone = $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.Phone)
		paramcount++
	}
	if filter.AgeFrom != nil {
		query += " and age >= $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.AgeFrom)
		paramcount++
	}
	if filter.AgeTo != nil {
		query += " and age <= $" + fmt.Sprintf("%d", paramcount)
		params = append(params, *filter.AgeTo)
		paramcount++
	}

	users := []models.User{}

	rows, err := u.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age,
			&user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

// read by one
func (u *UserRepo) GetUserById(userId string) (*models.User, error) {
	user := models.User{Id: userId}

	query := `select name, phone, age, created_at, updated_at
	from users where id = $1 and deleted_at is null`

	err := u.DB.QueryRow(query, userId).Scan(&user.Name, &user.Phone, &user.Age,
		&user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// update
func (u *UserRepo) UpdateUser(user *models.UpdateUser) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update users set name=$1, phone=$2, age=$3, updated_at=$4
	where id = $5 and deleted_at is null`
	result, err := tx.Exec(query, user.Name, user.Phone, user.Age,
		time.Now(), user.Id)

	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %v", user.Id, err)
	}
	return nil
}

// delete
func (u *UserRepo) DeleteUser(userID string) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update users set deleted_at = $1 where id = $2`
	result, err := tx.Exec(query, time.Now(), userID)
	if err != nil {
		return err
	}

	if numberOfAffectedRows, err := result.RowsAffected(); numberOfAffectedRows <= 0 {
		return fmt.Errorf("error: there is no row with the id = %s: %s", userID, err)
	}
	return nil
}
