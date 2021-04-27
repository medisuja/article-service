package usecase

import (
	HealthCheckInterface "article-service/app/health-check"
	"article-service/models"
)

type HealthCheckUsecase struct {
	HealthCheckRepository HealthCheckInterface.IHealthCheckRepository
}

func NewHealthCheckUsecase(h HealthCheckInterface.IHealthCheckRepository) HealthCheckInterface.IHealthCheckUsecase {
	return &HealthCheckUsecase{
		HealthCheckRepository: h,
	}
}

func (a *HealthCheckUsecase) GetDBTimestamp() models.HealthCheck {
	healthCheck := a.HealthCheckRepository.GetDBTimestamp()
	return healthCheck
}
