package svc

import (
	"log"

	"helloblog/internal/config"
	commentctrl "helloblog/internal/controller/comment"
	likectrl "helloblog/internal/controller/like"
	linkctrl "helloblog/internal/controller/link"
	notifctrl "helloblog/internal/controller/notification"
	postctrl "helloblog/internal/controller/post"
	searchctrl "helloblog/internal/controller/search"
	userctrl "helloblog/internal/controller/user"
	"helloblog/internal/dao"
	"helloblog/internal/pkg/jwt"
	commentservice "helloblog/internal/service/comment"
	likeservice "helloblog/internal/service/like"
	linkservice "helloblog/internal/service/link"
	notifservice "helloblog/internal/service/notification"
	postservice "helloblog/internal/service/post"
	searchservice "helloblog/internal/service/search"
	userservice "helloblog/internal/service/user"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	Redis       *redis.Client
	JWT         *jwt.Manager
	Post        postservice.UseCase
	Search      searchservice.UseCase
	Controllers Controllers
}

type Controllers struct {
	User         *userctrl.Controller
	Post         *postctrl.Controller
	Comment      *commentctrl.Controller
	Like         *likectrl.Controller
	Search       *searchctrl.Controller
	Notification *notifctrl.Controller
	Link         *linkctrl.Controller
}

// CommentNotifier 实现 commentservice.Notifier 接口
type CommentNotifier struct {
	postDAO      *dao.PostDAO
	notifService *notifservice.Service
	userDAO      *dao.UserDAO
}

func (n *CommentNotifier) NotifyComment(postID, fromUserID int64, content string) {
	post, err := n.postDAO.GetByID(postID)
	if err != nil || post == nil {
		return
	}
	if post.UserID == fromUserID {
		return // 自己的评论不通知
	}
	fromUser, err := n.userDAO.GetByID(fromUserID)
	if err != nil || fromUser == nil {
		return
	}
	title := fromUser.Username + " 评论了你的文章"
	text := content
	if len(text) > 100 {
		text = text[:100]
	}
	pid := postID
	fuid := fromUserID
	if err := n.notifService.Create(post.UserID, "comment", title, text, &fuid, &pid); err != nil {
		log.Printf("[notifier] create notification failed: %v", err)
	}
}

func NewServiceContext(cfg config.Config, database *gorm.DB, redisClient *redis.Client) *ServiceContext {
	jwtManager := jwt.NewManager(cfg.JWT.Secret, cfg.JWT.ExpireHours)

	// DAO
	userDAO := dao.NewUserDAO(database)
	postDAO := dao.NewPostDAO(database)
	commentDAO := dao.NewCommentDAO(database)
	likeDAO := dao.NewLikeDAO(database)
	searchDAO := dao.NewSearchDAO(database)
	notifDAO := dao.NewNotificationDAO(database)
	linkDAO := dao.NewLinkDAO(database)

	// Service
	userSvc := userservice.NewService(userDAO, jwtManager)
	postSvc := postservice.NewService(postDAO)
	notifSvc := notifservice.NewService(notifDAO)
	linkSvc := linkservice.NewService(linkDAO)
	notifier := &CommentNotifier{postDAO: postDAO, notifService: notifSvc, userDAO: userDAO}
	commentSvc := commentservice.NewService(commentDAO, postDAO, notifier)
	likeSvc := likeservice.NewService(likeDAO, postDAO, commentDAO)
	searchSvc := searchservice.NewService(searchDAO)

	// Controller
	userController := userctrl.NewController(userSvc)
	postController := postctrl.NewController(postSvc)
	commentController := commentctrl.NewController(commentSvc)
	likeController := likectrl.NewController(likeSvc)
	searchController := searchctrl.NewController(searchSvc)
	notifController := notifctrl.NewController(notifSvc)
	linkController := linkctrl.NewController(linkSvc)

	return &ServiceContext{
		Config: cfg,
		DB:     database,
		Redis:  redisClient,
		JWT:    jwtManager,
		Post:   postSvc,
		Search: searchSvc,
		Controllers: Controllers{
			User:         userController,
			Post:         postController,
			Comment:      commentController,
			Like:         likeController,
			Search:       searchController,
			Notification: notifController,
			Link:         linkController,
		},
	}
}
