package models

import (
	"errors"

	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/forms"
)

//User ...
type User struct {
	ID       int    `db:"user_id, primarykey, autoincrement" json:"user_id"`
	Name     string `db:"user_name" json:"user_name"`
	Email    string `db:"user_email" json:"user_email"`
	Password string `db:"user_password" json:"-"`
}

//UserModel ...
type UserModel struct {
}

//Signin ...
func (m UserModel) Signin(form forms.SigninForm) (user User, err error) {

	row := db.DB.QueryRow("SELECT user_id, user_name, user_email, user_password FROM tblUser WHERE user_email=$1 LIMIT 1", form.Email)
	var uid int
	var userName string
	var userEmail string
	var userPassword string

	err = row.Scan(&uid, &userName, &userEmail, &userPassword)
	if err != nil {
		return user, err
	}

	user.ID = uid
	user.Name = userName
	user.Email = userEmail
	user.Password = userPassword

	if form.Password == userPassword {
		return user, nil
	}
	return user, errors.New("Invalid password")
}

//GetOne ...
func (m UserModel) GetOne(userID int) (user User, err error) {
	row := db.DB.QueryRow("SELECT user_id, user_name, user_email, user_password FROM tblUser WHERE user_id=$1", userID)

	var uid int
	var userName string
	var userEmail string
	var userPassword string

	err = row.Scan(&uid, &userName, &userEmail, &userPassword)
	if err != nil {
		return User{}, err
	}

	user = User{uid, userName, userEmail, userPassword}

	return user, err
}
