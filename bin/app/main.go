package main

import (
	"fmt"
	"net/http"
	"runtime"

	"kumparan-backend-position-interview/bin/config"
	articles "kumparan-backend-position-interview/bin/modules/articles/handlers"
	"kumparan-backend-position-interview/bin/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	file := ""
	line := 0

	// Check the type of error and set appropriate status code and message
	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		message = e.Message.(string)
	default:
		// Get the file and line number where the error occurred
		if pc, file, line, ok := runtime.Caller(2); ok {
			message = fmt.Sprintf("Unexpected error at %s:%d: %v", file, line, err)

			// You can also access the function name using runtime.FuncForPC
			funcName := runtime.FuncForPC(pc).Name()
			message += fmt.Sprintf("\nFunction: %s", funcName)
		} else {
			message = fmt.Sprintf("Unexpected error: %v", err)
		}
	}

	// Send the error response
	c.JSON(code, map[string]interface{}{
		"error": message,
		"file":  file,
		"line":  line,
	})
}

func main() {
	e := echo.New()
	e.Validator = utils.NewValidationUtil()

	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Handle OPTIONS method for preflight requests
	e.OPTIONS("/*", func(c echo.Context) error {
		// Set the CORS headers in the response
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "*")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Access-Control-Max-Age", "86400") // Optional: Set cache duration in seconds

		return c.NoContent(http.StatusNoContent)
	})

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This service is running properly")
	})

	// Register the error handler middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					// Call the custom error handler function
					errorHandler(fmt.Errorf("%v", err), c)
				}
			}()
			return next(c)
		}
	})

	userGroup := e.Group("")

	articles.New().Mount(userGroup)

	listenerPort := fmt.Sprintf("localhost:%d", config.GlobalEnv.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}
