package models

import (
	"errors"
	"fmt"

	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/forms"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID       int    `db:"user_id, primarykey, autoincrement" json:"user_id"`
	Name     string `db:"user_name" json:"user_name"`
	Email    string `db:"user_email" json:"user_email"`
	Password string `db:"user_password" json:"-"`
	// UpdatedAt int  `db:"updated_at" json:"updated_at"`
	// CreatedAt int  `db:"created_at" json:"created_at"`
}

//UserModel ...
type UserModel struct{}

//Signin ...
func (m UserModel) Signin(form forms.SigninForm) (user User, err error) {

	row := db.DB.QueryRow("SELECT user_id, user_name, user_email, user_password FROM tblUser WHERE user_email=$1 LIMIT 1", form.Email)
	fmt.Println(form.Email)
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

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, errors.New("Invalid password")
	}

	fmt.Println(user.ID)

	return user, nil
}

//Signup ...
func (m UserModel) Signup(form forms.SignupForm) (user User, err error) {
	// getDb := db.GetDB()

	// checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

	// if err != nil {
	// 	return user, err
	// }

	// if checkUser > 0 {
	// 	return user, errors.New("User exists")
	// }

	// bytePassword := []byte(form.Password)
	// hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := getDb.Exec("INSERT INTO public.user(email, password, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())

	// if res != nil && err == nil {
	// 	err = getDb.SelectOne(&user, "SELECT id, email, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)
	// 	if err == nil {
	// 		return user, nil
	// 	}
	// }

	return user, errors.New("Not registered")
}

//One ...
func (m UserModel) One(userID int) (user User, err error) {
	// err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1", userID)
	return user, err
}

//GetOne ...
func (m UserModel) GetOne(userID int) (user User, err error) {
	// err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1", userID)

	fmt.Printf("GET ONE: %d \n", userID)

	// row := db.DBE.QueryRow("SELECT user_id, user_name, user_email, user_password FROM tblUser WHERE user_id=$1", userID)
	// dbaa := db.Init()
	row := db.DB.QueryRow("SELECT user_id, user_name, user_email, user_password FROM tblUser WHERE user_id=$1", userID)

	var uid int
	var userName string
	var userEmail string
	var userPassword string

	err = row.Scan(&uid, &userName, &userEmail, &userPassword)
	// err = row.Scan(&uid, &userName)

	fmt.Printf("WUT : %s \n", userName)

	if err != nil {
		return User{}, err
	}

	user = User{uid, userName, userEmail, userPassword}

	fmt.Printf("GOT USER: %+v \n", user)

	return user, err
}
