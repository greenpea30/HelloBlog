package post

import postservice "helloblog/internal/service/post"

type Controller struct {
	service postservice.UseCase
}

func NewController(service postservice.UseCase) *Controller {
	return &Controller{service: service}
}
