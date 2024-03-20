package service

import (
	"errors"

	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/repository"
)

type StudentService struct {
	StudentRepositry repository.StudentRepositryInterface
}

func NewStudentService(studentRepositry repository.StudentRepositryInterface) *StudentService {
	return &StudentService{
		StudentRepositry: studentRepositry,
	}
}

func (s *StudentService) CreateStudent(student *model.Student) error {
	if len(student.Sifre) < 5 {
		return errors.New("Şifre en az 5 karakter içermelidir")
	}
	return s.StudentRepositry.CreateStudent(student)

}

func (s *StudentService) GetStudentByID(id uint) (*model.Student, error) {
	return s.StudentRepositry.GetStudentByID(id)
}

func (s *StudentService) UpdateStudent(student *model.Student) error {
	return s.StudentRepositry.UpdateStudent(student)
}
func (s *StudentService) DeleteStudentById(id uint) error {
	return s.StudentRepositry.DeleteStudentById(id)
}
