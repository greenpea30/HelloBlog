package link

import (
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"
	linkservice "helloblog/internal/service/link"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service linkservice.UseCase
}

func NewController(service linkservice.UseCase) *Controller {
	return &Controller{service: service}
}

func (ctl *Controller) Create(c *gin.Context) {
	var req dto.CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
		return
	}
	resp, err := ctl.service.Create(req)
	if err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, resp)
}

func (ctl *Controller) List(c *gin.Context) {
	resp, err := ctl.service.List()
	if err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, resp)
}

func (ctl *Controller) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Fail(c, response.NewError(response.CodeInvalid, "invalid id"))
		return
	}
	if err := ctl.service.Delete(id); err != nil {
		response.Fail(c, response.Wrap(response.CodeInternalError, "internal server error", err))
		return
	}
	response.OK(c, nil)
}
