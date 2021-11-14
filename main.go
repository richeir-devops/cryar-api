package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID       uint
	Nickname string
	Age      uint8
}

func main() {
	fmt.Println("Hello Echo API")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users", getUsers)

	// Create a Consul API client
	// client, _ := api.NewClient(api.DefaultConfig())

	// Create an instance representing this service. "my-service" is the
	// name of _this_ service. The service should be cleaned up via Close.
	// svc, _ := connect.NewService("my-service", client)
	// defer svc.Close()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo API!")
}

func getUsers(c echo.Context) error {
	dsn := fmt.Sprintf("root:password@tcp(%s:3306)/db?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MySQL_IP_Address"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var user UserInfo
	db.Take(&user, 1)

	return c.String(http.StatusOK, user.Nickname)
}
