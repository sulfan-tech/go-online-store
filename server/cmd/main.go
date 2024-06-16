package main

import (
	"go-online-store/server/router"
	"log"
	"net/http"
	"os"

	valiator "go-online-store/internal/middleware/validator"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env := "../../.env"
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file:" + err.Error())
	}

	e := echo.New()

	// Register validator
	e.Validator = &valiator.CustomValidator{Validator: validator.New()}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost},
	}))

	e = router.RegisterRouter(e)
	e.Logger.Print("Server is running on port: " + os.Getenv("SERVER_PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
