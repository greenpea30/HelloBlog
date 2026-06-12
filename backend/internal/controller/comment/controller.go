package comment

import (
	commentservice "helloblog/internal/service/comment"

	//"github.com/gin-gonic/gin"
)

type Controller struct {
	service commentservice.UseCase
}

func NewController(service commentservice.UseCase) *Controller {
	return &Controller{service: service}
}
