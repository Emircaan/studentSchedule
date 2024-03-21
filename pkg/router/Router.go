package router

import (
	"github.com/emircaan/scheduleManager/pkg/controller"
	"github.com/labstack/echo/v4"
)

func SetupStudentRouters(e *echo.Echo, studentController controller.StudentControllerInterface) {

	e.POST("/students", studentController.CreateStudent)
	e.GET("/students/:id", studentController.GetStudentByID)
	e.DELETE("/students/:id", studentController.DeleteStudentById)
	e.PUT("/students/:id", studentController.UpdateStudent)
}

func SetupAuthRoutes(e *echo.Echo, authController *controller.AuthController) {
	e.POST("/login", authController.Login)
}
