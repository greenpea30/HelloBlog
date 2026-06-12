package user

import (
	"strings"

	"helloblog/internal/dao"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/password"
	"helloblog/internal/pkg/response"
)

func (s *Service) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.users.GetByEmail(email)
	if err != nil {
		if dao.IsNotFound(err) {
			return nil, response.NewError(response.CodeUnauthorized, "invalid email or password")
		}
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	if !password.Verify(user.PasswordHash, req.Password) {
		return nil, response.NewError(response.CodeUnauthorized, "invalid email or password")
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
