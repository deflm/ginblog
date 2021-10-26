package routes

import (
	"ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// user

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// category
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		// article
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		// 上传文件
		auth.POST("upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.POST("login", v1.Login)
		router.POST("user/add", v1.AddUser)
	}

	panic(r.Run(utils.HttpPort))
}
