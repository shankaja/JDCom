package main

import (
	"example/interview/config"
	"example/interview/util"
	"fmt"
)

func main() {
	cfg := config.LoadAll("config/config.yml")
	fmt.Print(*cfg)
	fmt.Print(util.ReadPdf(cfg.Files.Resume))
}
