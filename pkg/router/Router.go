package router

import (
	"net/http"

	"github.com/emircaan/scheduleManager/pkg/controller"
	"github.com/emircaan/scheduleManager/pkg/service"
	"github.com/labstack/echo/v4"
)

func SetupStudentRouters(e *echo.Echo, studentController controller.StudentControllerInterface) {

	e.POST("/students", studentController.CreateStudent)
	e.GET("/students/:id", studentController.GetStudentByID)
	e.DELETE("/students/:id", studentController.DeleteStudentById)
	e.PUT("/students/:id", studentController.UpdateStudent)
	e.GET("/students", studentController.GetStudents)

}

func SetupAuthRoutes(e *echo.Echo, authController *controller.AuthController) {
	e.POST("/login", authController.Login)
}

func SetupPlanRouters(e *echo.Echo, planController controller.IPlanController, authService service.AuthService) {
	r := e.Group("/plans", authMiddleware(authService))
	r.POST("", planController.CreatePlan)
	r.GET("", planController.GetPlans)
	r.GET("/:student_id", planController.GetPlansByStudentID)
	r.PUT("/:student_id/:plan_id", planController.UpdatePlanByStudentAndPlanID)
}

func authMiddleware(authService service.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token bulunamadı")
			}

			tokenString = tokenString[len("Bearer "):]

			isValid, err := authService.AuthenticateJWT(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token doğrulanırken bir hata oluştu")
			}
			if !isValid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Geçersiz token")
			}

			return next(c)
		}
	}
}
