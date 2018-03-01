//v0.0.2

package main

import (
/*  _ "github.com/go-sql-driver/mysql"
  "github.com/flosch/pongo2"*/
  "net/http"
  "AwsTest/models"
//  "fmt"
"github.com/labstack/echo"
  )

type User2 struct {
 *models.User
}


func main()  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!1")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
