package main

import (
	"bluewave-uptime-agent/internal/config"
	"bluewave-uptime-agent/internal/handler"
	"bluewave-uptime-agent/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := config.NewConfig(os.Getenv("PORT"), os.Getenv("API_SECRET"), os.Getenv("ALLOW_PUBLIC_API"))
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.AuthRequired(appConfig.ApiSecret))

	apiV1.GET("/health", handler.Health)
	apiV1.GET("/metrics", handler.Metrics)
	apiV1.GET("/ws", handler.WebSocket)

	server := &http.Server{
		Addr:    ":" + appConfig.Port,
		Handler: r.Handler(),
	}

	// TODO: Add graceful shutdown
	serve(server)

}

func serve(srv *http.Server) {
	srvErr := srv.ListenAndServe()
	if srvErr != nil && srvErr != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", srvErr)
	}
}
