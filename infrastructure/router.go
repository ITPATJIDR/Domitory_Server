package infrastructure

import (
	"Domitory_Server/database"
	"Domitory_Server/interfaces"

	"github.com/gin-gonic/gin"
)

func Dispatch(s database.GormDB) {
	r := gin.Default()
	userController := interfaces.NewUserController(s)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", userController.Register)
	r.POST("/getUserByID", userController.GetUserByID)

	r.Run()
}
