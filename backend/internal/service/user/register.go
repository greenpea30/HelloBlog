package user

import (
	"strings"

	//"helloblog/internal/dao"
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/password"
	"helloblog/internal/pkg/response"
)

func (s *Service) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if exists, err := s.users.UsernameExists(req.Username); err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	} else if exists {
		return nil, response.NewError(response.CodeConflict, "username already taken")
	}

	if req.Email != "" {
		if exists, err := s.users.EmailExists(req.Email); err != nil {
			return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
		} else if exists {
			return nil, response.NewError(response.CodeConflict, "email already registered")
		}
	}

	hash, err := password.Hash(req.Password)
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: hash,
	}
	if req.Email != "" {
		user.Email = &req.Email
	}

	if err := s.users.Create(user); err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	token, err := s.jwt.Generate(user.ID)
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	return &dto.AuthResponse{
		User:        toUserResponse(user),
		AccessToken: token,
	}, nil
}
