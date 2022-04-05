package main

import (
	"example/interview/config"
	"fmt"
)

func main() {
	cfg := config.LoadAll("config/config.yml")
	fmt.Print(*cfg)
}
