package controllers

import (
	"strconv"

	"github.com/ndbeals/britannicus-final-project/forms"
	"github.com/ndbeals/britannicus-final-project/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct {
}

var userModel = new(models.UserModel)

//getUserID ...
func GetUserID(c *gin.Context) int {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		return models.ConvertToInt64(userID)
	}
	return 0
}

//GetLoggedinUser
func GetLoggedinUser(c *gin.Context) (user models.User, success bool) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID != nil {
		user.ID = models.ConvertToInt64(userID)
		user.Name = session.Get("user_name").(string)
		user.Email = session.Get("user_email").(string)
	} else {
		return user, false
	}

	return user, true
}

//Signin ...
func (ctrl UserController) Signin(c *gin.Context) {
	var signinForm forms.SigninForm
	var success bool

	signinForm.Email, success = c.GetPostForm("email")
	if !success {
		c.Abort()
	}
	signinForm.Password, success = c.GetPostForm("password")
	if !success {
		c.Abort()
	}

	user, err := userModel.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)

		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Name)
		session.Save()

		// c.IndentedJSON(200, gin.H{"message": "User signed in", "user": user})
		c.Redirect(303, "/")
	} else {
		c.IndentedJSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
	}
}

//Signout ...
func (ctrl UserController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(303, "/")
}

//GetOne ...
func (ctrl UserController) GetOne(c *gin.Context) {
	userid := c.Param("id")

	if userid, err := strconv.ParseInt(userid, 10, 32); err == nil {
		userid := int(userid)

		data, err := userModel.GetOne(userid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.IndentedJSON(200, gin.H{"data": data})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
