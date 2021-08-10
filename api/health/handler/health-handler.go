package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/prometheus-converter/api/domain"

	log "github.com/sirupsen/logrus"
)

type HealthHandler struct {
	HealthUsecase domain.HealthUsecase
}

func NewHealthHandler(c *gin.Engine, healthUsecase domain.HealthUsecase) {
	handler := &HealthHandler{
		HealthUsecase: healthUsecase,
	}

	api := c.Group("api/v1")
	{
		api.GET("prometheus", handler.GetPrometheusHealth)
	}
}

func (s *HealthHandler) GetPrometheusHealth(c *gin.Context) {

	log.Printf("HealthHandler GetPrometheusHealth...")

	metrics, _ := s.HealthUsecase.GetPrometheusHealthInfo(c)
	c.Header("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

	c.Data(http.StatusOK, "text/plain; version=0.0.4; charset=utf-8", []byte(metrics))
}
