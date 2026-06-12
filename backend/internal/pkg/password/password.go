package password

import "golang.org/x/crypto/bcrypt"

const cost = 12

// Hash 对密码进行哈希
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Verify 校验密码
func Verify(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashOnlyForZJU 为 ZJU 用户生成占位密码哈希（不用于密码登录）
func HashOnlyForZJU(studentID string) string {
	hash, _ := Hash("zju:" + studentID)
	return hash
}
