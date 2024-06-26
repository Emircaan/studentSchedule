package repository

import "github.com/emircaan/scheduleManager/pkg/model"

type StudentRepositryInterface interface {
	CreateStudent(student *model.Student) error
	GetStudents() ([]model.Student, error)
	GetStudentByID(id uint) (*model.Student, error)
	UpdateStudent(student *model.Student) error
	DeleteStudentById(id uint) error
	GetStudentByEmail(email string) (*model.Student, error)
}
