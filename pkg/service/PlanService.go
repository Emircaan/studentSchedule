package service

import (
	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/repository"
)

type PlanService struct {
	PlanRepository repository.IPlanRepository
}

func NewPlanService(planRepository repository.IPlanRepository) IPlanService {
	return &PlanService{
		PlanRepository: planRepository,
	}
}

func (s *PlanService) CreatePlan(plan *model.Plan) error {
	return s.PlanRepository.CreatePlan(plan)
}

func (s *PlanService) GetPlans() ([]model.Plan, error) {
	return s.PlanRepository.GetPlans()
}

func (s *PlanService) UpdatePlanByStudentAndPlanID(studentID uint, planID uint, plan *model.Plan) error {
	return s.PlanRepository.UpdatePlanByStudentAndPlanID(studentID, planID, plan)
}

func (s *PlanService) GetPlansByStudentID(studentID uint) ([]model.Plan, error) {
	return s.PlanRepository.GetPlansByStudentID(studentID)

}
