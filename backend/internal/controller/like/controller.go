package like

import (
	likeservice "helloblog/internal/service/like"

	//"github.com/gin-gonic/gin"
)

type Controller struct {
	service likeservice.UseCase
}

func NewController(service likeservice.UseCase) *Controller {
	return &Controller{service: service}
}
