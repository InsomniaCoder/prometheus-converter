package handler

import (
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

	// if err != nil {
	// 	log.Panicf("store sensor data fail: %+v\n", err)
	// } else {
	// 	c.JSON(200, savedSensor)
	// }

}
