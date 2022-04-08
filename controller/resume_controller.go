package controller

import (
	"example/interview/config"
	"example/interview/service"
	"example/interview/wrap"
	"log"
	"net/http"
)

type ResumeController struct {
	c             *config.Config
	resumeService *service.ResumeService
}

func NewResumeController(c *config.Config, rs *service.ResumeService) *ResumeController {
	return &ResumeController{c: c, resumeService: rs}
}

func (c *ResumeController) UploadResume(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 1)
	file, handler, err := r.FormFile("resume")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	text, err := c.resumeService.Upload(file, handler)
	if err != nil {
		log.Fatal(err)
		return
	}

	wrap.WrapFileResponse(w, text)
	return
}

func (c *ResumeController) CompareResume(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 1)
	file, handler, err := r.FormFile("resume")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	text, err := c.resumeService.Upload(file, handler)
	if err != nil {
		log.Fatal(err)
		return
	}

	wrap.WrapFileResponse(w, text)
	return
}
