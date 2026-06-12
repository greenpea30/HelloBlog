package notification

import (
	notificationservice "helloblog/internal/service/notification"
)

type Controller struct {
	service notificationservice.UseCase
}

func NewController(service notificationservice.UseCase) *Controller {
	return &Controller{service: service}
}
