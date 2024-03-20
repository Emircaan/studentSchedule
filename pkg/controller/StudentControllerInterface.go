package controller

import "github.com/labstack/echo/v4"

type StudentControllerInterface interface {
	CreateStudent(ctx echo.Context) error
	GetStudentByID(ctx echo.Context) error
	UpdateStudent(ctx echo.Context) error
	DeleteStudentById(ctx echo.Context) error
}
