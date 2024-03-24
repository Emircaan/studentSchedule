package repository

import (
	"github.com/emircaan/scheduleManager/pkg/model"

	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{
		db: db,
	}
}

func (r *PlanRepository) CreatePlan(plan *model.Plan) error {
	return r.db.Create(plan).Error
}

func (r *PlanRepository) GetPlans() ([]model.Plan, error) {
	var plans []model.Plan
	err := r.db.Find(&plans).Error
	return plans, err
}

func (r *PlanRepository) UpdatePlanByStudentAndPlanID(studentID uint, planID uint, plan *model.Plan) error {
	return r.db.Model(&model.Plan{}).Where("student_id = ? AND id = ?", studentID, planID).Updates(plan).Error
}

func (r *PlanRepository) GetPlansByStudentID(studentID uint) ([]model.Plan, error) {
	var plans []model.Plan
	err := r.db.Where("student_id = ?", studentID).Find(&plans).Error
	return plans, err
}
