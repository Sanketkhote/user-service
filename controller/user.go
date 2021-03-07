package controller

import (
	"fmt"
	"net/http"

	"github.com/Sanketkhote/microService/service/user"
	"github.com/Sanketkhote/microService/service/user/model"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	SaveUser(c *gin.Context)
}
type controller struct {
	user user.User
}

func NewController() Controller {
	return &controller{
		user: user.NewUser(),
	}

}

func (c *controller) SaveUser(ctx *gin.Context) {
	var user model.UserModel
	ctx.BindJSON(&user)
	fmt.Println(user)
	if user.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	exist, err := c.user.SaveUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exist {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Already Present"})
		return
	}
	ctx.String(http.StatusCreated, "Created")

}
