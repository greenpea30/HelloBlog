package like

import (
	"strconv"

	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type ToggleRequest struct {
	TargetType string `json:"target_type" binding:"required,oneof=post comment"`
	TargetID   int64  `json:"target_id" binding:"required,min=1"`
}

func (ctl *Controller) Toggle(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req ToggleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	liked, err := ctl.service.Toggle(userID, req.TargetType, req.TargetID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, gin.H{
		"target_type": req.TargetType,
		"target_id":   req.TargetID,
		"liked":       liked,
	})
}

func (ctl *Controller) UserLikedPosts(c *gin.Context) {
	userID := c.GetInt64("user_id")
	ids, err := ctl.service.GetUserLikedPostIDs(userID)
	if err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, gin.H{"liked_post_ids": ids})
}

// 辅助函数，从路径参数中解析 id
func parseIDParam(c *gin.Context, name string) (int64, error) {
	idStr := c.Param(name)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, response.NewError(response.CodeInvalid, "invalid "+name)
	}
	return id, nil
}
