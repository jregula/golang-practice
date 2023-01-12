package routes

import (
    "example/web-service-gin/controllers" //add this
    "github.com/gin-gonic/gin"
)
func UserRoute(router *gin.Engine)  {
    //All routes related to users comes here
    router.POST("/user", controllers.CreateUser()) //add this
}