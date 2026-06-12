package user

import userservice "helloblog/internal/service/user"

type Controller struct {
	service userservice.UseCase
}

func NewController(service userservice.UseCase) *Controller {
	return &Controller{service: service}
}
