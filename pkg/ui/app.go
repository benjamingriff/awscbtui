package ui

import (
	"fmt"
)

type Config struct {
	Name string
}

func Run(cfg *Config) error {
	fmt.Printf("Welcome to %s\n", cfg.Name)
	return nil
}
