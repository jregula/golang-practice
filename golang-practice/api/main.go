// https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m
package main

import (
        "example/web-service-gin/configs" //add this
        "github.com/gin-gonic/gin"
        "example/web-service-gin/routes" //add this

    )
    
func main() {
        router := gin.Default()
  
        configs.ConnectDB()

        routes.UserRoute(router) //add this

        router.Run()
}