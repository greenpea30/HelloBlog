package server

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"helloblog/internal/server/middleware"
	"helloblog/internal/svc"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, svcCtx *svc.ServiceContext) {
	// 静态文件服务（上传的图片）
	engine.Static("/uploads", "./uploads")

	// 文件上传
	api := engine.Group("/api/v1")

	uploadRoutes := api.Group("/upload")
	uploadRoutes.Use(middleware.Auth(svcCtx.JWT))
	{
		uploadRoutes.POST("/avatar", handleAvatarUpload)
	}

	registerUserRoutes(api, svcCtx)
	registerPostRoutes(api, svcCtx)
	registerCommentRoutes(api, svcCtx)
	registerLikeRoutes(api, svcCtx)
	registerSearchRoutes(api, svcCtx)
	registerNotificationRoutes(api, svcCtx)
	registerLinkRoutes(api, svcCtx)
	registerFolderRoutes(api, svcCtx)
}

func registerUserRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", svcCtx.Controllers.User.Register)
		auth.POST("/login", svcCtx.Controllers.User.Login)
		auth.POST("/zju-login", svcCtx.Controllers.User.ZJULogin)
	}

	userRoutes := api.Group("/users")
	userRoutes.Use(middleware.Auth(svcCtx.JWT))
	{
		userRoutes.GET("/me", svcCtx.Controllers.User.Me)
		userRoutes.PUT("/me", svcCtx.Controllers.User.UpdateProfile)
		userRoutes.GET("/me/posts", svcCtx.Controllers.Post.ListMyPosts)
	}
}

func registerPostRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	posts := api.Group("/posts")
	{
		posts.GET("", svcCtx.Controllers.Post.List)
		posts.GET("/:id", svcCtx.Controllers.Post.GetByID)

		authPosts := posts.Group("")
		authPosts.Use(middleware.Auth(svcCtx.JWT))
		{
			authPosts.POST("", svcCtx.Controllers.Post.Create)
			authPosts.PUT("/:id", svcCtx.Controllers.Post.Update)
			authPosts.DELETE("/:id", svcCtx.Controllers.Post.Delete)
		}
	}
}

func registerCommentRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	api.GET("/posts/:id/comments", svcCtx.Controllers.Comment.ListByPost)

	authComments := api.Group("/comments")
	authComments.Use(middleware.Auth(svcCtx.JWT))
	{
		authComments.DELETE("/:id", svcCtx.Controllers.Comment.Delete)
	}

	api.POST("/posts/:id/comments", middleware.Auth(svcCtx.JWT), svcCtx.Controllers.Comment.Create)
}

func registerLikeRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	likes := api.Group("/likes")
	likes.Use(middleware.Auth(svcCtx.JWT))
	{
		likes.POST("/toggle", svcCtx.Controllers.Like.Toggle)
		likes.GET("/user-liked-posts", svcCtx.Controllers.Like.UserLikedPosts)
	}
}

func registerNotificationRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	notifs := api.Group("/notifications")
	notifs.Use(middleware.Auth(svcCtx.JWT))
	{
		notifs.GET("", svcCtx.Controllers.Notification.List)
		notifs.GET("/unread-count", svcCtx.Controllers.Notification.UnreadCount)
		notifs.POST("/mark-all-read", svcCtx.Controllers.Notification.MarkAllRead)
	}
}

func registerLinkRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	api.GET("/links", svcCtx.Controllers.Link.List)
	links := api.Group("/links")
	links.Use(middleware.Auth(svcCtx.JWT))
	{
		links.POST("", svcCtx.Controllers.Link.Create)
		links.DELETE("/:id", svcCtx.Controllers.Link.Delete)
	}
}

func registerFolderRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	// 公开：用户主页（无需登录）
	api.GET("/users/:id/profile", svcCtx.Controllers.Folder.GetUserProfile)

	// 需要登录
	folders := api.Group("/folders")
	folders.Use(middleware.Auth(svcCtx.JWT))
	{
		folders.GET("", svcCtx.Controllers.Folder.List)
		folders.POST("", svcCtx.Controllers.Folder.Create)
		folders.DELETE("/:id", svcCtx.Controllers.Folder.Delete)
	}
}

func handleAvatarUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "请选择文件"})
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".png"
	}
	filename := fmt.Sprintf("avatar_%d%s", time.Now().UnixNano(), ext)
	outPath := filepath.Join("uploads", filename)
	out, err := os.Create(outPath)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "保存文件失败"})
		return
	}
	defer out.Close()
	io.Copy(out, file)

	c.JSON(200, gin.H{"code": 0, "msg": "success", "data": gin.H{"url": "/uploads/" + filename}})
}

func registerSearchRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	api.GET("/search", svcCtx.Controllers.Search.FullTextSearch)
}
