package servicecontainer

import (
	"example/interview/config"
	"example/interview/controller"
	"example/interview/service"
)

type DependancyContainer struct {
	config *config.Config

	// controller
	HealthController *controller.HealthController
	ResumeController *controller.ResumeController

	// service
	ResumeService *service.ResumeService
}

func (d *DependancyContainer) Init() *config.Config {
	d.config = config.LoadAll("config/config.yml")

	d.ResumeService = service.NewResumeService(d.config)

	d.HealthController = controller.NewHealthController(d.config)
	d.ResumeController = controller.NewResumeController(d.config, d.ResumeService)

	return d.config
}
