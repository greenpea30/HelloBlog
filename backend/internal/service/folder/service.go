package folder

import (
	"helloblog/internal/dao"
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"

	"gorm.io/gorm"
)

type UseCase interface {
	Create(userID int64, req dto.CreateFolderRequest) (*dto.FolderResponse, error)
	List(userID int64) ([]dto.FolderResponse, error)
	Delete(id, userID int64) error
	GetUserProfile(userID int64, viewerID int64) (*dto.UserProfileResponse, error)
}

type folderRepository interface {
	Create(folder *model.Folder) error
	GetByID(id int64) (*model.Folder, error)
	ListByUser(userID int64) ([]model.Folder, error)
	Delete(id int64, userID int64) error
}

type postRepository interface {
	ListByUserAndFolder(userID int64, folderID *int64) ([]model.Post, error)
	CountByUserAndFolder(userID int64, folderID *int64) (int64, error)
}

type userRepository interface {
	GetByID(id int64) (*model.User, error)
}

type Service struct {
	folders folderRepository
	posts   postRepository
	users   userRepository
	db      *gorm.DB
}

func NewService(db *gorm.DB, folders folderRepository, posts postRepository, users userRepository) *Service {
	return &Service{db: db, folders: folders, posts: posts, users: users}
}

func (s *Service) Create(userID int64, req dto.CreateFolderRequest) (*dto.FolderResponse, error) {
	folder := &model.Folder{
		UserID: userID,
		Name:   req.Name,
	}
	if err := s.folders.Create(folder); err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}
	return &dto.FolderResponse{
		ID:        folder.ID,
		Name:      folder.Name,
		PostCount: 0,
		CreatedAt: folder.CreatedAt,
	}, nil
}

func (s *Service) List(userID int64) ([]dto.FolderResponse, error) {
	folders, err := s.folders.ListByUser(userID)
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}
	resp := make([]dto.FolderResponse, 0, len(folders)+1)
	// 未分类
	uncategorizedCount, _ := s.posts.CountByUserAndFolder(userID, nil)
	resp = append(resp, dto.FolderResponse{
		ID:        0,
		Name:      "未分类",
		PostCount: int(uncategorizedCount),
	})
	for _, f := range folders {
		count, _ := s.posts.CountByUserAndFolder(userID, &f.ID)
		resp = append(resp, dto.FolderResponse{
			ID:        f.ID,
			Name:      f.Name,
			PostCount: int(count),
			CreatedAt: f.CreatedAt,
		})
	}
	return resp, nil
}

func (s *Service) Delete(id, userID int64) error {
	// 先清空该文件夹下文章的 folder_id
	if err := dao.ClearFolderID(s.db, id); err != nil {
		return response.Wrap(response.CodeInternalError, "internal server error", err)
	}
	return s.folders.Delete(id, userID)
}

func (s *Service) GetUserProfile(userID int64, viewerID int64) (*dto.UserProfileResponse, error) {
	user, err := s.users.GetByID(userID)
	if err != nil {
		if dao.IsNotFound(err) {
			return nil, response.NewError(response.CodeNotFound, "user not found")
		}
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	folders, err := s.folders.ListByUser(userID)
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	profile := &dto.UserProfileResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			AvatarURL: user.AvatarURL,
			Bio:       user.Bio,
		},
		Folders: make([]dto.FolderWithPosts, 0),
	}

	// 未分类
	uncategorizedPosts, _ := s.posts.ListByUserAndFolder(userID, nil)
	if len(uncategorizedPosts) > 0 {
		fwp := dto.FolderWithPosts{
			ID:    0,
			Name:  "未分类",
			Posts: make([]dto.PostResponse, len(uncategorizedPosts)),
		}
		for i, p := range uncategorizedPosts {
			fwp.Posts[i] = toPostResponse(&p)
		}
		profile.Folders = append(profile.Folders, fwp)
	}

	// 各文件夹
	for _, f := range folders {
		posts, _ := s.posts.ListByUserAndFolder(userID, &f.ID)
		if len(posts) == 0 {
			continue
		}
		fwp := dto.FolderWithPosts{
			ID:    f.ID,
			Name:  f.Name,
			Posts: make([]dto.PostResponse, len(posts)),
		}
		for i, p := range posts {
			fwp.Posts[i] = toPostResponse(&p)
		}
		profile.Folders = append(profile.Folders, fwp)
	}

	return profile, nil
}

func toPostResponse(post *model.Post) dto.PostResponse {
	format := post.Format
	if format == "" {
		format = "markdown"
	}
	return dto.PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		Summary:      post.Summary,
		Format:       format,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		ViewCount:    post.ViewCount,
		CreatedAt:    post.CreatedAt,
	}
}
