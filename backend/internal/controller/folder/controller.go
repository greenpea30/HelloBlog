package folder

import folderservice "helloblog/internal/service/folder"

type Controller struct {
	service folderservice.UseCase
}

func NewController(service folderservice.UseCase) *Controller {
	return &Controller{service: service}
}
