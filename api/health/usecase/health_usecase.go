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

func (s *healthUsecase) GetPrometheusHealthInfo(c context.Context) (string, error) {
	log.Println("healthUsecase GetPrometheusHealthInfo...")

	res, err := s.healthRepository.GetHealthInfo(c)

	var prometheus_stat_text string

	if err != nil {
		log.Errorf("call API error %v\n", err)
		prometheus_stat_text += "gateway_up 0"
		return prometheus_stat_text, nil
	}

	prometheus_stat_text += "gateway_up 1\n"

	var face_comparison_stat string
	for _, instance := range res.FaceComparison.Instances {
		face_comparison_stat = "face_comparison_up 0\n"
		if instance.Status == "ok" {
			face_comparison_stat = "face_comparison_up 1\n"
			break
		}
	}
	prometheus_stat_text += face_comparison_stat

	var thai_id_stat string
	for _, instance := range res.ThaiID.Instances {
		thai_id_stat = "thai_id_up 0\n"
		if instance.Status == "ok" {
			thai_id_stat = "thai_id_up 1\n"
			break
		}
	}
	prometheus_stat_text += thai_id_stat

	var antispoofing_stat string
	for _, instance := range res.Antispoofing.Instances {
		antispoofing_stat = "antispoofing_up 0\n"
		if instance.Status == "ok" {
			antispoofing_stat = "antispoofing_up 1\n"
			break
		}
	}
	prometheus_stat_text += antispoofing_stat

	return prometheus_stat_text, nil
}
