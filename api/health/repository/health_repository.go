package repository

import (
	"context"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/insomniacoder/prometheus-converter/api/domain"
)

type healthRepository struct {
	targetAPI string
}

func NewHealthRepository(target string) domain.HealthRepository {
	return &healthRepository{targetAPI: target}
}

func (s *healthRepository) GetHealthInfo(c context.Context) (info *domain.HealthInfo, err error) {
	log.Println("healthRepository GetHealthInfo...")

	response, err := http.Get(s.targetAPI)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer response.Body.Close()

	var healthInfo domain.HealthInfo

	json.NewDecoder(response.Body).Decode(&healthInfo)

	log.Println("health body: %v", healthInfo)

	return &healthInfo, err
}
