package router

import (
	"bluebell1/controller"
	"bluebell1/logger"
	"bluebell1/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册业务
	r.POST("/signup", controller.SignUpHandler)
	//登录业务
	r.POST("/login", controller.LoginHandler)

	r.Use(middlewares.JWTAuthMiddleware())
	{
		r.GET("/community", controller.CommunityHandler)
		r.GET("/community/:id", controller.CommunityDetailHandler)
		r.POST("/post", controller.CreatePostHandler)
		r.GET("/post/:id", controller.PostDetailHandler)
		r.GET("/posts", controller.GetPostListHandler)
		r.POST("/vote", controller.PostVoteController)
		r.GET("/posts2", controller.GetPostListHandler2)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
