package healthcheck

import "article-service/models"

type IHealthCheckUsecase interface {
	GetDBTimestamp() models.HealthCheck
}
