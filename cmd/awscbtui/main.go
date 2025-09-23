package main

import (
	"github.com/benjamingriff/awscbtui/pkg/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
