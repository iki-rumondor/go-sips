package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/utils"
	"gorm.io/gorm"
)

type UserService struct {
	Repo interfaces.UserRepoInterface
}

func NewUserService(repo interfaces.UserRepoInterface) interfaces.UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) VerifyUser(req *request.SignIn) (string, error) {

	user, err := s.Repo.FindUserBy("username", req.Username)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &response.Error{
				Code:    401,
				Message: "Username atau Password Salah",
			}
		}
		return "", response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return "", &response.Error{
			Code:    401,
			Message: "Username atau password salah",
		}
	}

	jwt, err := utils.GenerateToken(user.Uuid, user.Role.Name)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *UserService) GetCountSubjects() (*response.SubjectsCount, error) {
	subjects, err := s.Repo.FindSubjects()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	practicalSubjects, err := s.Repo.FindPracticalSubjects()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return &response.SubjectsCount{
		General:   len(*subjects),
		Practical: len(*practicalSubjects),
	}, nil

}
