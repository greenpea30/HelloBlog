package comment

import (
	"strconv"

	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")

	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || postID <= 0 {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid post id"))
		return
	}

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.Create(userID, postID, req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) ListByPost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || postID <= 0 {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid post id"))
		return
	}

	resp, err := ctl.service.ListByPost(postID)
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
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid comment id"))
		return
	}

	if err := ctl.service.Delete(id, userID); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, nil)
}
