package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/repository"
)

type AuthService struct {
	UserRepository repository.StudentRepositryInterface
}

func NewAuthService(userRepository repository.StudentRepositryInterface) AuthService {
	return AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Authentication(email, password string) (*model.Student, error) {

	user, err := s.UserRepository.GetStudentByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil || user.Sifre != password {
		return nil, errors.New("Kullanıcı adı veya şifre yanlış")
	}

	return user, nil

}
func (s *AuthService) GenerateJWT(user *model.Student) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
