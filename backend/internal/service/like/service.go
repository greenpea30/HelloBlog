package like

// UseCase 点赞接口
type UseCase interface {
	Toggle(userID int64, targetType string, targetID int64) (liked bool, err error)
	GetUserLikedPostIDs(userID int64) ([]int64, error)
}

// LikeCounter 点赞计数更新接口
type LikeCounter interface {
	IncrementLikeCount(id int64) error
	DecrementLikeCount(id int64) error
}

type likeRepository interface {
	Toggle(userID int64, targetType string, targetID int64) (liked bool, err error)
	GetUserLikedPostIDs(userID int64) ([]int64, error)
}

type Service struct {
	likes          likeRepository
	postCounter    LikeCounter
	commentCounter LikeCounter
}

func NewService(likes likeRepository, postCounter LikeCounter, commentCounter LikeCounter) *Service {
	return &Service{
		likes:          likes,
		postCounter:    postCounter,
		commentCounter: commentCounter,
	}
}

func (s *Service) GetUserLikedPostIDs(userID int64) ([]int64, error) {
	return s.likes.GetUserLikedPostIDs(userID)
}
