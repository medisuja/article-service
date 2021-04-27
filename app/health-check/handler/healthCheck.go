package handler

import (
	"time"

	HealthCheckInterface "article-service/app/health-check"

	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	HealthStatus string    `json:"healthStatus"`
	DBTimestamp  time.Time `json:"databaseTimestamp"`
}

type HealthCheckHandler struct {
	HealthCheckUsecase HealthCheckInterface.IHealthCheckUsecase
}

func (a *HealthCheckHandler) Check(c *gin.Context) {
	healthCheck := a.HealthCheckUsecase.GetDBTimestamp()
	res := &HealthCheckResponse{
		HealthStatus: "GOOD",
		DBTimestamp:  healthCheck.CurrentTimestamp,
	}
	c.JSON(200, res)
}
