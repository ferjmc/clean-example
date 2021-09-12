package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	v1 := e.Group("/api/v1")

	health := v1.Group("/health")

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
