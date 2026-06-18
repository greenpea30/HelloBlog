package post

import (
	"helloblog/internal/dao"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"
)

func (s *Service) Create(userID int64, req dto.CreatePostRequest) (*dto.PostResponse, error) {
	post, err := s.posts.Create(toModel(userID, req))
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}
	resp := toResponse(post)
	return &resp, nil
}

func (s *Service) GetByID(id int64) (*dto.PostResponse, error) {
	post, err := s.posts.GetByID(id)
	if err != nil {
		return nil, response.Wrap(response.CodeNotFound, "post not found", err)
	}

	incd, _ := s.posts.IncrementView(id)
	if incd {
		post.ViewCount++
	}

	resp := toResponse(post)
	return &resp, nil
}

func (s *Service) Update(id int64, userID int64, req dto.UpdatePostRequest) (*dto.PostResponse, error) {
	post, err := s.posts.GetByID(id)
	if err != nil {
		return nil, response.Wrap(response.CodeNotFound, "post not found", err)
	}

	if post.UserID != userID {
		return nil, response.NewError(response.CodeForbidden, "not your post")
	}

	post.Title = req.Title
	post.Summary = req.Summary
	post.Content = req.Content
	if req.Format != "" {
		post.Format = req.Format
	}

	if err := s.posts.Update(post); err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	resp := toResponse(post)
	return &resp, nil
}

func (s *Service) Delete(id int64, userID int64) error {
	return s.posts.SoftDelete(id, userID)
}

func (s *Service) List(req dto.PostListRequest) (*dto.PostListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	params := dao.PostListParams{
		Page:     req.Page,
		PageSize: req.PageSize,
		UserID:   req.UserID,
		OrderBy:  req.OrderBy,
		ZJUOnly:  req.ZJUOnly,
	}

	posts, total, err := s.posts.List(params)
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	items := make([]dto.PostResponse, len(posts))
	for i, p := range posts {
		items[i] = toResponse(&p)
	}

	totalPages := (total + int64(req.PageSize) - 1) / int64(req.PageSize)

	return &dto.PostListResponse{
		Items:      items,
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
	}, nil
}
