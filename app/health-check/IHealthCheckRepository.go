package healthcheck

import "article-service/models"

type IHealthCheckRepository interface {
	GetDBTimestamp() models.HealthCheck
}
