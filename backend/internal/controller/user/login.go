package user

import (
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.Login(req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}
