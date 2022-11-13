package models

import (
	"log"
	"time"
)

type User struct {
	ID         int
	UUID       string
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}

func (u *User) CreateUser() (err error) { // user作成
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) { // user取得(単体)
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Created_at,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) { // user更新
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) { // user削除
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
