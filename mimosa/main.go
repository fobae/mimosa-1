package main

import (
    "github.com/gin-gonic/gin"
)

func index (c *gin.Context) {
    id := c.Query("id")
    content := gin.H{"Hello": id}
    c.JSON(200, content)
}

func main() {
  app := gin.Default()
  app.GET("/", index)
  app.Run(":1337")
}
      