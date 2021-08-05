package usecase

import (
	"context"

	"github.com/insomniacoder/prometheus-converter/api/domain"
	log "github.com/sirupsen/logrus"
)

type healthUsecase struct {
	healthRepository domain.HealthRepository
}

func NewHealthUsecase(healthRepo domain.HealthRepository) domain.HealthUsecase {
	return &healthUsecase{
		healthRepository: healthRepo,
	}
}

func (s *healthUsecase) GetPrometheusHealthInfo(ctx context.Context) (string, error) {
	log.Println("healthUsecase GetPrometheusHealthInfo...")

	return "", nil
}
