package controller

import "github.com/labstack/echo/v4"

type IPlanController interface {
	CreatePlan(ctx echo.Context) error
	GetPlans(ctx echo.Context) error
	UpdatePlanByStudentAndPlanID(ctx echo.Context) error
	GetPlansByStudentID(ctx echo.Context) error
}
