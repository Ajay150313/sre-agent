package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ajay150313/sre-agent/internal/agent"
	"github.com/ajay150313/sre-agent/internal/config"
)

func main() {
	cfg := config.Load()

	if cfg.OpenAIKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	analyzer := agent.NewAnalyzer(cfg.PrometheusURL, cfg.OpenAIKey)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				incidents, err := analyzer.AnalyzeAllAlerts(ctx)
				if err != nil {
					log.Printf("Error analyzing alerts: %v", err)
					continue
				}
				log.Printf("Analyzed %d alerts", len(incidents))
			case <-ctx.Done():
				return
			}
		}
	}()

	r := gin.Default()
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "SRE Agent",
		})
	})

	r.GET("/api/incidents", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"incidents": []interface{}{}})
	})

	r.POST("/api/analyze", func(c *gin.Context) {
		incidents, err := analyzer.AnalyzeAllAlerts(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, incidents)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
