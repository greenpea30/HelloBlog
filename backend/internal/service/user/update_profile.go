package user

import (
	"strings"

	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"
)

func (s *Service) UpdateProfile(userID int64, req dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	user, err := s.users.GetByID(userID)
	if err != nil {
		return nil, response.Wrap(response.CodeNotFound, "user not found", err)
	}

	req.Username = strings.TrimSpace(req.Username)

	if req.Username != user.Username {
		if exists, err := s.users.UsernameExists(req.Username); err != nil {
			return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
		} else if exists {
			return nil, response.NewError(response.CodeConflict, "username already taken")
		}
	}

	user.Username = req.Username
	user.AvatarURL = strings.TrimSpace(req.AvatarURL)
	user.Bio = strings.TrimSpace(req.Bio)

	if err := s.users.Update(user); err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	resp := toUserResponse(user)
	return &resp, nil
}
