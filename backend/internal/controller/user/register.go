package user

import (
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.Register(req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) ZJULogin(c *gin.Context) {
	var req dto.ZJULoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.ZJULogin(req.StudentID, req.Password)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}
