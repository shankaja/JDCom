package controller

import (
	"example/interview/config"
	"log"
	"net/http"
)

type HealthController struct {
	c *config.Config
}

func NewHealthController(c *config.Config) *HealthController {
	return &HealthController{c: c}
}

func (h *HealthController) GetHealth(w http.ResponseWriter, r *http.Request) {

	log.Printf("Health check with ctx")
	w.Write([]byte("Up and running"))
}
