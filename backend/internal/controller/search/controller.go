package search

import (
	searchservice "helloblog/internal/service/search"

	//"github.com/gin-gonic/gin"
)

type Controller struct {
	service searchservice.UseCase
}

func NewController(service searchservice.UseCase) *Controller {
	return &Controller{service: service}
}
