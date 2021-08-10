package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/prometheus-converter/config"

	_healthHandler "github.com/insomniacoder/prometheus-converter/api/health/handler"
	_healthRepository "github.com/insomniacoder/prometheus-converter/api/health/repository"
	_healthUsecase "github.com/insomniacoder/prometheus-converter/api/health/usecase"

	log "github.com/sirupsen/logrus"
)

func init() {

	log.Println("Initializting API package")
	log.Printf("Loaded configuration: %v", config.Config)
	//set up router
	r := gin.Default()
	r.Use(Cors())
	r.Static("/static", "./static")
	//set up health dependency
	healthRepo := _healthRepository.NewHealthRepository(config.Config.Target.URL)
	healthUsecase := _healthUsecase.NewHealthUsecase(healthRepo)
	_healthHandler.NewHealthHandler(r, healthUsecase)

	//start server
	portNumber := fmt.Sprintf(":%d", config.Config.Server.Port)
	r.Run(portNumber)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
