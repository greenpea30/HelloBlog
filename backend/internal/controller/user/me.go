package user

import (
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Me(c *gin.Context) {
	userID := c.GetInt64("user_id")

	resp, err := ctl.service.GetMe(userID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}

func (ctl *Controller) UpdateProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}

	resp, err := ctl.service.UpdateProfile(userID, req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.OK(c, resp)
}
