package post

import (
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req dto.CreatePostRequest
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

func (ctl *Controller) GetByID(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		response.Fail(c, err)
		return
	}

	resp, err := ctl.service.GetByID(id)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) Update(c *gin.Context) {
	userID := c.GetInt64("user_id")

	id, err := parseIDParam(c)
	if err != nil {
		response.Fail(c, err)
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.Update(id, userID, req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) Delete(c *gin.Context) {
	userID := c.GetInt64("user_id")

	id, err := parseIDParam(c)
	if err != nil {
		response.Fail(c, err)
		return
	}

	if err := ctl.service.Delete(id, userID); err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, nil)
}

func (ctl *Controller) ListMyPosts(c *gin.Context) {
	userID := c.GetInt64("user_id")
	var req dto.PostListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}
	req.UserID = userID
	resp, err := ctl.service.List(req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, resp)
}

func (ctl *Controller) List(c *gin.Context) {
	var req dto.PostListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.List(req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}
