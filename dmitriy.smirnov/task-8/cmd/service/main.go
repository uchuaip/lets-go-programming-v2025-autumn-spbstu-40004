package main

import (
	"fmt"

	"uchuaip/task-8/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Print(cfg.Environment, " ", cfg.LogLevel)
}
