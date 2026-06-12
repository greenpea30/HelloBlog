package link

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

type UseCase interface {
	Create(req dto.CreateLinkRequest) (*dto.LinkResponse, error)
	List() ([]dto.LinkResponse, error)
	Delete(id int64) error
}

type linkRepository interface {
	Create(link *model.Link) error
	List() ([]model.Link, error)
	Delete(id int64) error
}

type Service struct {
	repo linkRepository
}

func NewService(repo linkRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	link := &model.Link{Name: req.Name, URL: req.URL}
	if err := s.repo.Create(link); err != nil {
		return nil, err
	}
	return &dto.LinkResponse{ID: link.ID, Name: link.Name, URL: link.URL, CreatedAt: link.CreatedAt}, nil
}

func (s *Service) List() ([]dto.LinkResponse, error) {
	links, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	resp := make([]dto.LinkResponse, len(links))
	for i, l := range links {
		resp[i] = dto.LinkResponse{ID: l.ID, Name: l.Name, URL: l.URL, CreatedAt: l.CreatedAt}
	}
	return resp, nil
}

func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
