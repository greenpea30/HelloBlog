package user

import (
	"helloblog/internal/dao"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"
)

func (s *Service) GetMe(userID int64) (*dto.UserResponse, error) {
	user, err := s.users.GetByID(userID)
	if err != nil {
		if dao.IsNotFound(err) {
			return nil, response.NewError(response.CodeNotFound, "user not found")
		}
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	resp := toUserResponse(user)
	return &resp, nil
}
