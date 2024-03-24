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

func (s AuthService) AuthenticateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Geçersiz imza metodu")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("Geçersiz token")
	}

	return true, nil
}
