package app

import (
  "strconv"
  "log"

  "github.com/gin-gonic/gin"

  "finproc/app/models"
  "finproc/app/routes"
)

// StartServer Wraps the gin Router and binds some Middleware
func StartServer(ctx models.AppContext) {

  router := gin.New()

  // Logger middleware will write the logs to gin.DefaultWriter
  router.Use(gin.Logger())

  // Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

  routes.Routing(router, ctx)

  var port = strconv.Itoa(ctx.Port)
  log.Println("App started. Listening on port " + port + " in " + ctx.Env + " mode.")
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
