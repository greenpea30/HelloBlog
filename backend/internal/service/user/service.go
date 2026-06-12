package user

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
	"helloblog/internal/pkg/jwt"
)

// UseCase 定义用户模块接口（依赖倒置）
type UseCase interface {
	Register(req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req dto.LoginRequest) (*dto.AuthResponse, error)
	GetMe(userID int64) (*dto.UserResponse, error)
	ZJULogin(studentID, password string) (*dto.AuthResponse, error)
	UpdateProfile(userID int64, req dto.UpdateProfileRequest) (*dto.UserResponse, error)
}

// userRepository DAO 依赖接口
type userRepository interface {
	Create(user *model.User) error
	GetByID(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetByZjuID(zjuID string) (*model.User, error)
	UsernameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
	ZjuIDExists(zjuID string) (bool, error)
	Update(user *model.User) error
}

type Service struct {
	users userRepository
	jwt   *jwt.Manager
}

func NewService(users userRepository, jwtManager *jwt.Manager) *Service {
	return &Service{users: users, jwt: jwtManager}
}
