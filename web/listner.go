package web

import (
	"github.com/Sanketkhote/microService/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller controller.Controller
}

func NewRouter() *Router {
	return &Router{
		controller: controller.NewController(),
	}
}
func (r *Router) StartListner() {

	router := gin.Default()
	router.POST("/user",r.controller.SaveUser)
	router.Run(":8080")
}
