package repository

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/insomniacoder/prometheus-converter/api/domain"
)

type healthRepository struct {
}

func NewHealthRepository() domain.HealthRepository {
	return &healthRepository{}
}

func (s *healthRepository) GetHealthInfo(ctx context.Context) (info *domain.HealthInfo, err error) {
	log.Println("healthRepository GetHealthInfo...")
	return nil, nil
}
