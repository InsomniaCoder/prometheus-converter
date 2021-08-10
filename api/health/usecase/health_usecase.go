package usecase

import (
	"context"
	"fmt"

	"github.com/insomniacoder/prometheus-converter/api/domain"
	log "github.com/sirupsen/logrus"
)

type prometheusHealthMetric struct {
	gatewayUp      int
	facecompareUp  int
	thiaIdUp       int
	antiSpoofingUp int
}

type healthUsecase struct {
	healthRepository domain.HealthRepository
}

func NewHealthUsecase(healthRepo domain.HealthRepository) domain.HealthUsecase {
	return &healthUsecase{
		healthRepository: healthRepo,
	}
}

func (s *healthUsecase) GetPrometheusHealthInfo(c context.Context) (string, error) {
	log.Println(`healthUsecase GetPrometheusHealthInfo...`)

	res, err := s.healthRepository.GetHealthInfo(c)

	var metric = &prometheusHealthMetric{}

	if err != nil {

		log.Errorf(`call API error %v\n`, err)
		return formatPrometheusText(&prometheusHealthMetric{
			gatewayUp:      1,
			facecompareUp:  0,
			thiaIdUp:       0,
			antiSpoofingUp: 0,
		}), nil
	}

	metric.gatewayUp = 1

	for _, instance := range res.FaceComparison.Instances {
		metric.facecompareUp = 0
		if instance.Status == `ok` {
			metric.facecompareUp = 1
			break
		}
	}

	for _, instance := range res.ThaiID.Instances {
		metric.thiaIdUp = 0
		if instance.Status == `ok` {
			metric.thiaIdUp = 1
			break
		}
	}

	for _, instance := range res.Antispoofing.Instances {
		metric.antiSpoofingUp = 0
		if instance.Status == `ok` {
			metric.antiSpoofingUp = 1
			break
		}
	}
	return formatPrometheusText(metric), nil
}

func formatPrometheusText(m *prometheusHealthMetric) string {
	var promText = `gateway_up %d
face_comparison_up %d
thai_id_up %d
antispoofing_up %d
`
	return fmt.Sprintf(promText, m.gatewayUp, m.facecompareUp, m.thiaIdUp, m.antiSpoofingUp)
}
