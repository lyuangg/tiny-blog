package models

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	ID        int
	Name      string
	Token     string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func UserLogin(email string, password string) User {
	user := User{}
	password = EncryptPwd(password)

	Db.QueryRow("select id, user_name, user_email, user_token from tiny_users where user_email=? and user_password=?", email, password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Token)

	if user.ID > 0 {
		tokenStr := strconv.Itoa(user.ID) + "_" + time.Now().String()
		token := fmt.Sprintf("%x", md5.Sum([]byte(tokenStr)))
		err := UpdateUserToken(user.ID, token)
		if err == nil {
			user.Token = token
		}
	}

	return user
}

func UpdateUserToken(id int, token string) error {
	stmt, err := Db.Prepare("update tiny_users set user_token = ?, user_updated_at =? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(token, time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserToken(id int, token string) User {
	user := User{}
	Db.QueryRow("select id, user_name, user_email, user_token from tiny_users where id=? and user_token=?",
		id, token).Scan(&user.ID, &user.Name, &user.Email, &user.Token)
	return user
}

func UpdateUserName(id int, name string) error {
	stmt, err := Db.Prepare("update tiny_users set user_name= ?, user_updated_at =? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPwd(id int, password string) error {
	password = EncryptPwd(password)
	stmt, err := Db.Prepare("update tiny_users set user_password= ?, user_updated_at =? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(password, time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return err
	}
	return nil
}

func EncryptPwd(password string) string {
	password += "yuancoder"
	data := md5.Sum([]byte(password))
	md5Pwd := fmt.Sprintf("%x", data)

	return md5Pwd
}
