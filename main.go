package main

import (
	"fmt"
	_ "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	e := echo.New()
	db := connect()
	fmt.Println(db)

	e.GET("/", hello)
	e.POST("/student", create)
	e.DELETE("/student/:id", delete)
	e.GET("/students", getAll)
	e.PUT("/student/:id", Update)
	e.PATCH("/student/:id", Patch)

	e.Logger.Fatal(e.Start(":1323"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
