package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"github.com/emircaan/scheduleManager/pkg/controller"
	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/repository"
	"github.com/emircaan/scheduleManager/pkg/router"
	"github.com/emircaan/scheduleManager/pkg/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
)

type App struct {
	Router     *echo.Echo
	DB         *gorm.DB
	Controller controller.StudentControllerInterface
}

func main() {

	a := &App{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	a.initilaize(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	model.Migrate(a.DB)
	if err != nil {
		log.Fatal("Tablolar oluşturulamadı: ", err)
	}
	studentRepository := repository.NewStudentRepository(a.DB)
	studentService := service.NewStudentService(studentRepository)
	a.Controller = controller.NewStudentController(studentService)
	authService := service.NewAuthService(studentRepository)
	authController := controller.NewAuthController(authService)
	router.SetupAuthRoutes(a.Router, authController)
	planRepository := repository.NewPlanRepository(a.DB)
	planService := service.NewPlanService(planRepository)
	planController := controller.NewPlanController(planService)
	router.SetupPlanRouters(a.Router, planController, authService)

	router.SetupStudentRouters(a.Router, a.Controller)

	a.Router.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))

}

func (a *App) initilaize(username, password, host, port, dbname string) {
	var err error
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	a.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {

		log.Fatal(err)
	}
	a.Router = echo.New()

}
