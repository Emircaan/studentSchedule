package service

import (
	"github.com/emircaan/scheduleManager/pkg/model"
)

type IPlanService interface {
	CreatePlan(plan *model.Plan) error
	GetPlans() ([]model.Plan, error)
	UpdatePlanByStudentAndPlanID(studentID uint, planID uint, plan *model.Plan) error
	GetPlansByStudentID(studentID uint) ([]model.Plan, error)
}
