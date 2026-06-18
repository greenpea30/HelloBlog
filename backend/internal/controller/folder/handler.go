package folder

import (
	"strconv"

	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req dto.CreateFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.Create(userID, req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) List(c *gin.Context) {
	userID := c.GetInt64("user_id")

	resp, err := ctl.service.List(userID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) Delete(c *gin.Context) {
	userID := c.GetInt64("user_id")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid id"))
		return
	}

	if err := ctl.service.Delete(id, userID); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, nil)
}

func (ctl *Controller) GetUserProfile(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || userID <= 0 {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid user id"))
		return
	}

	// 可选：当前登录用户 ID（用于后续判断是否是自己）
	viewerID, _ := c.Get("user_id")
	_ = viewerID

	resp, err := ctl.service.GetUserProfile(userID, 0)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}
