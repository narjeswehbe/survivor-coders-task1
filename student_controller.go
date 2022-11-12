package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/entity"
	"net/http"
)

func create(c echo.Context) error {
	gormDB := connect()
	u := new(entity.Student)
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	fmt.Println(u)
	result := gormDB.Create(&u)
	fmt.Println(result.Row())
	gormDB.Save(&u)
	return c.JSON(http.StatusCreated, "Student created Successfully")
}
func delete(c echo.Context) error {
	id := c.Param("id")
	gormDB := connect()
	deleteMe := gormDB.Delete(&entity.Student{}, id)
	if deleteMe == nil {
		return c.JSON(http.StatusNotFound, "student not found")
	}
	fmt.Println(deleteMe)
	return c.JSON(http.StatusAccepted, "deleted successfully")
}
func getAll(c echo.Context) error {

	var Students []*entity.Student
	gormDB := connect()
	if err := gormDB.Find(&Students).Error; err != nil {
		// error handling here
		return err
	}

	return c.JSON(http.StatusAccepted, Students)

}
func Update(c echo.Context) error {

	gormDB := connect()
	id := c.Param("id")
	student := new(entity.Student)
	gormDB.First(&student, id)
	u := new(entity.Student)
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	if u.Firstname == "" {
		return c.JSON(http.StatusBadRequest, "firstname is required")
	}
	if u.Lastname == "" {
		return c.JSON(http.StatusBadRequest, "lastname is required")
	}

	student.Firstname = u.Firstname
	student.Lastname = u.Lastname
	gormDB.Save(student)
	return c.JSON(http.StatusAccepted, student)

}
func Patch(c echo.Context) error {

	gormDB := connect()
	id := c.Param("id")
	student := new(entity.Student)
	gormDB.First(&student, id)
	u := new(entity.Student)
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	if u.Firstname != "" {
		student.Firstname = u.Firstname
	}
	if u.Lastname != "" {
		student.Lastname = u.Lastname
	}
	gormDB.Save(student)
	return c.JSON(http.StatusAccepted, student)

}
