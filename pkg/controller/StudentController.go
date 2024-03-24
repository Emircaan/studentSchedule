package controller

import (
	"net/http"
	"strconv"

	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/service"
	"github.com/labstack/echo/v4"
)

type StudentController struct {
	StudentService service.StudentServiceInterface
}

func NewStudentController(studentService service.StudentServiceInterface) StudentControllerInterface {
	return &StudentController{
		StudentService: studentService,
	}
}

func (c *StudentController) CreateStudent(ctx echo.Context) error {
	student := new(model.Student)
	if err := ctx.Bind(student); err != nil {
		return err
	}

	if err := c.StudentService.CreateStudent(student); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, student)
}

func (c *StudentController) GetStudentByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	student, err := c.StudentService.GetStudentByID(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, student.ToResponse())
}

func (c *StudentController) DeleteStudentById(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.StudentService.DeleteStudentById(uint(id))
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *StudentController) UpdateStudent(ctx echo.Context) error {
	student := new(model.Student)
	if err := ctx.Bind(student); err != nil {
		return err
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	existingStudent, err := c.StudentService.GetStudentByID(uint(id))
	if err != nil {
		return err
	}

	existingStudent.Ad = student.Ad
	existingStudent.Soyad = student.Soyad
	existingStudent.Eposta = student.Eposta
	existingStudent.Sifre = student.Sifre

	if err := c.StudentService.UpdateStudent(existingStudent); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, existingStudent)
}

func (c *StudentController) GetStudentByEmail(ctx echo.Context) error {
	email := ctx.Param("email")
	student, err := c.StudentService.GetStudentByEmail(email)
	if err != nil {
		return err
	} else if student == nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) GetStudents(ctx echo.Context) error {
	students := make([]map[string]string, 0)

	ss, err := c.StudentService.GetStudents()
	if err != nil {

		return err
	}
	for _, s := range ss {
		students = append(students, s.ToResponse())
	}
	return ctx.JSON(http.StatusOK, students)
}
