package controller

import (
	"net/http"
	"strconv"

	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/service"
	"github.com/labstack/echo/v4"
)

type PlanController struct {
	PlanService service.IPlanService
}

func NewPlanController(planService service.IPlanService) IPlanController {
	return &PlanController{
		PlanService: planService,
	}
}

func (s *PlanController) CreatePlan(ctx echo.Context) error {
	plan := new(model.Plan)
	if err := ctx.Bind(plan); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, s.PlanService.CreatePlan(plan))
}

func (s *PlanController) GetPlans(ctx echo.Context) error {
	var plans []model.Plan

	plans, err := s.PlanService.GetPlans()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, plans)

}

func (s *PlanController) UpdatePlanByStudentAndPlanID(ctx echo.Context) error {
	studentIDStr := ctx.Param("student_id")
	planIDStr := ctx.Param("plan_id")

	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		return err
	}

	planID, err := strconv.ParseUint(planIDStr, 10, 64)
	if err != nil {
		return err
	}

	var plan model.Plan
	if err := ctx.Bind(&plan); err != nil {
		return err
	}

	if err := s.PlanService.UpdatePlanByStudentAndPlanID(uint(studentID), uint(planID), &plan); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "Plan güncellendi")
}

func (s *PlanController) GetPlansByStudentID(ctx echo.Context) error {
	studentID := ctx.Param("student_id")
	studentIDUint, err := strconv.ParseUint(studentID, 10, 64)
	if err != nil {
		return err
	}

	plans, err := s.PlanService.GetPlansByStudentID(uint(studentIDUint))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, plans)
}
