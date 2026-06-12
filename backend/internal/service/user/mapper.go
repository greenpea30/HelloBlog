package user

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

func toUserResponse(user *model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		ZjuID:     user.ZjuID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
	}
}
