package post

import (
	"helloblog/internal/dao"
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

// UseCase 定义文章模块接口
type UseCase interface {
	Create(userID int64, req dto.CreatePostRequest) (*dto.PostResponse, error)
	GetByID(id int64) (*dto.PostResponse, error)
	Update(id int64, userID int64, req dto.UpdatePostRequest) (*dto.PostResponse, error)
	Delete(id int64, userID int64) error
	List(params dto.PostListRequest) (*dto.PostListResponse, error)
}

// postRepository DAO 依赖接口
type postRepository interface {
	Create(post *model.Post) (*model.Post, error)
	GetByID(id int64) (*model.Post, error)
	Update(post *model.Post) error
	SoftDelete(id int64, userID int64) error
	List(params dao.PostListParams) ([]model.Post, int64, error)
	IncrementView(id int64) (bool, error)
}

type Service struct {
	posts postRepository
}

func NewService(posts postRepository) *Service {
	return &Service{posts: posts}
}
