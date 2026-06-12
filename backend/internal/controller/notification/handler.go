package notification

import (
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) List(c *gin.Context) {
	userID := c.GetInt64("user_id")
	resp, err := ctl.service.List(userID)
	if err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, resp)
}

func (ctl *Controller) UnreadCount(c *gin.Context) {
	userID := c.GetInt64("user_id")
	count, err := ctl.service.UnreadCount(userID)
	if err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, gin.H{"unread_count": count})
}

func (ctl *Controller) MarkAllRead(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if err := ctl.service.MarkAllRead(userID); err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, nil)
}
