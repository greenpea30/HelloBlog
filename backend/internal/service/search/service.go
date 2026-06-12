package search

import (
	"helloblog/internal/dto"
)

type UseCase interface {
	FullTextSearch(req dto.SearchRequest) (*dto.SearchResponse, error)
}

type searchRepository interface {
	FullTextSearch(query string, limit int) ([]dto.SearchResultItem, error)
}

type Service struct {
	repo searchRepository
}

func NewService(repo searchRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FullTextSearch(req dto.SearchRequest) (*dto.SearchResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	items, err := s.repo.FullTextSearch(req.Query, req.PageSize)
	if err != nil {
		return nil, err
	}

	return &dto.SearchResponse{
		Query:    req.Query,
		Items:    items,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
