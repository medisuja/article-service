package repository

import (
	HealthCheckInterface "article-service/app/health-check"
	"article-service/models"

	"github.com/jinzhu/gorm"
)

type HealthCheckRepository struct {
	Conn *gorm.DB
}

var (
	HTTPPROXY = "" //config.Config.KAI.HttpProxy
	PORTPROXY = "" // config.Config.KAI.PortProxy
)

func NewHealthCheckRepository(Conn *gorm.DB) HealthCheckInterface.IHealthCheckRepository {
	return &HealthCheckRepository{Conn}
}

func (m *HealthCheckRepository) GetDBTimestamp() models.HealthCheck {
	var healthCheck models.HealthCheck

	tx := m.Conn.Begin()
	tx.Raw("SELECT current_timestamp").Scan(&healthCheck)
	tx.Commit()

	return healthCheck
}
