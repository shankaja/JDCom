package main

import (
	"example/interview/config"
	"example/interview/controller"
	"example/interview/route"
	"fmt"
)

func main() {
	cfg := config.LoadAll("config/config.yml")
	fmt.Print(*cfg)
	healthController := controller.NewHealthController(cfg)
	resumeController := controller.NewResumeController(cfg)
	route.NewRouter(healthController, resumeController, cfg).InitializeRouter()
}
