package api

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/healthz"),
		gin.Recovery(),
	)
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/api/v1/healthz", Healthz)
	router.GET("/", Flip)

	return router
}

func StartRouter() {
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		log.Debugf("API_PORT not set, defaulting to 8080.")
		apiPort = "8080"
	}

	log.Infof("Starting API on port %s", apiPort)

	r := SetupRouter()

	err := r.Run(":" + apiPort)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
