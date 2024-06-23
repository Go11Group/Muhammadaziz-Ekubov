package postgres

import (
	"Muhammadaziz-Ekubov/3-moth-homework/3-homework/model"
	"database/sql"
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

		rows.Scan(&user.ID)
	}

}
