package route

import (
	"example/interview/config"
	"example/interview/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	healthController *controller.HealthController
	resumeController *controller.ResumeController
	c                *config.Config
}

func NewRouter(hc *controller.HealthController, rc *controller.ResumeController, c *config.Config) *Router {
	return &Router{healthController: hc, c: c, resumeController: rc}
}

func (t *Router) InitializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/health", t.healthController.GetHealth)
	r.HandleFunc("/upload", t.resumeController.UploadResume)
	r.HandleFunc("/compare", t.resumeController.UploadResume)
	log.Fatal(http.ListenAndServe(":8000", r))
}
