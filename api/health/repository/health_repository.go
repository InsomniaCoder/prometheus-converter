package repository

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/insomniacoder/prometheus-converter/api/domain"
)

type healthRepository struct {
	targetAPI string
}

func NewHealthRepository(target string) domain.HealthRepository {
	return &healthRepository{targetAPI: target}
}

func (s *healthRepository) GetHealthInfo(ctx context.Context) (info *domain.HealthInfo, err error) {
	log.Println("healthRepository GetHealthInfo...")
	return nil, nil
}
