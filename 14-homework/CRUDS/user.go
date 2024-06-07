package CRUDS

import (
	"Go11/14-homework/model"
	"Go11/14-homework/storage/postgres"
	"database/sql"
	"encoding/json"
	"net/http"
)

type CrudUser struct {
	db *sql.DB
}

func NewCrudUsersRepo(db *sql.DB) *CrudUser{
	return &CrudUser{db: db}
}

func (i *CrudUser) InsertUsers(w http.ResponseWriter, r *http.Request) {
	user, err := postgres.InsertUser(w, r)
	if err != nil {
		panic(err)
	}

	_, err = i.db.Exec("insert into users(id,firstname,lastname,age) values($1,$2,$3,$4)",
		&user.Id, &user.FirstName, &user.LastName, &user.Age)
	if err != nil {
		panic(err)
	}
}

func (u *CrudUser) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	up, err := postgres.UpdateUser(w, r)
	if err != nil {
		panic(err)
	}

	_, err = u.db.Exec("update users set firstname=$1, lastname=$2,age=$3 where id=$4",
		&up.FirstName, &up.LastName, &up.Age, &up.Id)

	if err != nil {
		panic(err)
	}
}

func (d *CrudUser) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	del, err := postgres.DeleteUser(w, r)
	if err != nil {
		panic(err)
	}

	_, err = d.db.Exec("delete from users where id=$1", &del.Id)
	if err != nil {
		panic(err)
	}
}

func (s *CrudUser) ReadUsers(w http.ResponseWriter, r *http.Request) {
	row, err := s.db.Query("select * from users")
	if err != nil {
		http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var users []model.Users
	for row.Next() {
		var user model.Users
		err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age)
		if err != nil {
			http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err)
	}
}
