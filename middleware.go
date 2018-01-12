package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndbeals/britannicus-final-project/controllers"
)

// authentication middleware
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := controllers.GetUserID(c)

		if userID == 0 {
			c.Redirect(303, "/")
		}

		c.Next()

	}
}
