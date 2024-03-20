package repository

import (
	"github.com/emircaan/scheduleManager/pkg/model"
	"gorm.io/gorm"
)

type StudentRepositry struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepositry {
	return &StudentRepositry{
		DB: db,
	}
}

func (r *StudentRepositry) CreateStudent(student *model.Student) error {
	return r.DB.Create(student).Error
}

func (r *StudentRepositry) GetStudentByID(id uint) (*model.Student, error) {
	var student model.Student
	err := r.DB.First(&student, id).Error
	return &student, err
}

func (r *StudentRepositry) UpdateStudent(student *model.Student) error {
	return r.DB.Save(student).Error
}

func (r *StudentRepositry) DeleteStudentById(id uint) error {
	return r.DB.Delete(&model.Student{}, id).Error
}
