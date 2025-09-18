package main

import (
	"github.com/benjamingriff/awscbtui/pkg/ui"
)

func main() {
	cfg := &ui.Config{
		Name: "awscbtui",
	}

	if err := ui.Run(cfg); err != nil {
		panic(err)
	}
}
