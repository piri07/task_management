package main

import (
    "github.com/gin-gonic/gin"
    "task_management/config"
    "task_management/routes"
)

func main() {
    r := gin.Default()
    config.ConnectDatabase()
    routes.RegisterTaskRoutes(r)
    r.Run(":8080")
}
