package user

import (
	"context"
	"strings"

	"helloblog/internal/dao"
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"
	"helloblog/internal/zjulogin"

	pwd "helloblog/internal/pkg/password"
)

func (s *Service) ZJULogin(studentID, zjuPassword string) (*dto.AuthResponse, error) {
	studentID = strings.TrimSpace(studentID)

	// 验证 ZJU 学号密码
	cfg := zjulogin.Config{Username: studentID, Password: zjuPassword}
	auth, err := zjulogin.New(cfg)
	if err != nil {
		return nil, response.NewError(response.CodeUnauthorized, "学号或密码错误: "+err.Error())
	}
	if err := auth.ZJUAM().Login(context.Background()); err != nil {
		return nil, response.NewError(response.CodeUnauthorized, err.Error())
	}

	// 查找或创建用户
	user, err := s.users.GetByZjuID(studentID)
	if err != nil && !dao.IsNotFound(err) {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	if user == nil {
		// 新用户：用学号作为用户名（如果已存在则追加数字）
		username := "zju_" + studentID
		if exists, _ := s.users.UsernameExists(username); exists {
			username = "zju_" + studentID + "_1"
		}
		user = &model.User{
			ZjuID:        &studentID,
			Username:     username,
			PasswordHash: pwd.HashOnlyForZJU(studentID),
		}
		if err := s.users.Create(user); err != nil {
			return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
		}
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
