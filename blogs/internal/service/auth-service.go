package service

import (
	"github.com/Quanghh2233/blogs/internal/dto"
	"github.com/Quanghh2233/blogs/internal/model"
	"github.com/Quanghh2233/blogs/internal/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(userDto dto.User) (*gorm.DB, model.User)
	VerifyCredential(email string, password string) (bool, uint64)
}

type authService struct {
	authRepo repository.AuhtRepo
}

func NewAuthService(authRepo repository.AuhtRepo) *authService {
	return &authService{authRepo: authRepo}
}

func (service *authService) Register(userDTO dto.User) (*gorm.DB, model.User) {
	userModel := model.User{}
	err := smapping.FillStruct(&userModel, smapping.MapFields(&userDTO))
	if err != nil {
		panic(err)
	}
	return service.authRepo.Register(userModel)
}

func (service *authService) VerifyCredential(email string, password string) (bool, uint64) {
	result, user := service.authRepo.FindByEmail(email)
	if result.Error == nil && user.ID != 0 {
		return comparePassword([]byte(user.Password), []byte(password)), user.ID
	}
	return false, 0
}

func comparePassword(hashedPass []byte, plainPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, plainPass)
	return err == nil
}
