package main

import (
	"example/interview/route"
	"example/interview/servicecontainer"
)

func main() {

	dc := &servicecontainer.DependancyContainer{}
	cfg := dc.Init()

	route.NewRouter(dc.HealthController, dc.ResumeController, cfg).InitializeRouter()
}
