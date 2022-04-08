package service

import (
	"example/interview/config"
	"example/interview/util"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type ResumeService struct {
	config *config.Config
}

func NewResumeService(c *config.Config) *ResumeService {
	return &ResumeService{config: c}
}

func (s *ResumeService) Upload(file multipart.File, handler *multipart.FileHeader) (string, error) {
	log.Printf("Uploaded File: %+v\n", handler.Filename)

	tempFile, err := ioutil.TempFile("temp", "upload-*.pdf")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	tempFile.Write(fileBytes)
	text, err := util.ReadPdf(fmt.Sprintf("%+v", tempFile.Name()))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return text, nil
}

//Reads the resume.pdf text and job description
//Returns the similarity percentage of those two
func (s *ResumeService) Compare(file multipart.File, handler *multipart.FileHeader, jd string) (string, error) {
	text, err := s.Upload(file, handler)
	if err != nil {
		return "", err
	}
	//TODO integrate the NLP engine

	return text, nil
}
