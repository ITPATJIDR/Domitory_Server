package interfaces

import (
	"Domitory_Server/database"
	"Domitory_Server/domain"
	"Domitory_Server/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
}

func NewUserController(s database.GormDB) *UserController {
	return &UserController{
		UserInteractor: usecases.UserInteractor{
			DB: s,
		},
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}

	status, err := uc.UserInteractor.FindByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	if status {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already registered"})
		return
	}

	hashedPassword, err := usecases.GenerateFromPassword([]byte(user.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	user.Password = string(hashedPassword)

	if err := uc.UserInteractor.NewUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered Successfully"})
}

func (uc *UserController) GetUserByID(c *gin.Context) {

	var userID domain.UserID

	if err := c.Bind(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	user, err := uc.UserInteractor.FindByID(userID.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
