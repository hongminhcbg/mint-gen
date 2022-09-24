package router

import (
	"github.com/gin-gonic/gin"
  "sms/src/service"
)

func InitGin(e *gin.Engine, s *service.Service) {
	e.POST("/api/v1/users", s.CreateUser)
}
