package controller

import (
	"example/interview/config"
	"example/interview/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResumeController struct {
	c *config.Config
}

func NewResumeController(c *config.Config) *ResumeController {
	return &ResumeController{c: c}
}

func (c *ResumeController) UploadResume(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 1)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)

	tempFile, err := ioutil.TempFile("temp", "upload-*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	text, error := util.ReadPdf(fmt.Sprintf("%+v", tempFile.Name()))
	if error != nil {
		log.Print(err)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File  %+s \n", text)
}
