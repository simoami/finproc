package routes

import (
  "github.com/gin-gonic/gin"

  "finproc/app/handlers"
  "finproc/app/models"
)

func Routing(router *gin.Engine, ctx models.AppContext) {
  
  // Route definitions
  router.GET("/healthcheck", makeHandler(ctx, handlers.Healthcheck))
}

//  Closure function to add the app context to route handlers
func makeHandler(ctx models.AppContext, fn func(*gin.Context, models.AppContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, ctx)
	}
}
