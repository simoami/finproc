package handlers

import (
  "github.com/gin-gonic/gin"

  "finproc/app/models"
)

func Healthcheck(c *gin.Context, ctx models.AppContext) {

  c.JSON(200, gin.H{
    "status": "OK",
  })
}
