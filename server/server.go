package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)


// StockAPI struct
type StockAPI struct {
	e *echo.Echo
}

// NewServer Instance of Echo
func NewServer() *StockAPI {
	return &StockAPI{
		e: echo.New(),
	}
}

// Start server

func (s *StockAPI) Start(port string) {
	// logger
	s.e.Use(middleware.Logger())
	// recover
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.DELETE, echo.PATCH, echo.POST},
	}))
	// price endpoint
	//s.e.GET("/price", )

	s.e.GET("/", func(c  echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})

	// Start Server
	s.e.Logger.Fatal(s.e.Start(port))

}

// Close  server functionality
func (s *StockAPI) Close()  {
	s.e.Close()
}

