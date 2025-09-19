package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("pipelines", 0, 0, maxX/2-1, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Pipelines:")
		fmt.Fprintln(v, "- Pipeline 1")
		fmt.Fprintln(v, "- Pipeline 2")
		fmt.Fprintln(v, "- Pipeline 3")
	}

	if v, err := g.SetView("details", maxX/2, 0, maxX-1, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Pipeline Details:")
		fmt.Fprintln(v, "Status: Running")
		fmt.Fprintln(v, "Last Build: Success")
	}

	if v, err := g.SetView("status", 0, maxY-4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Status: Connected to AWS | Press Ctrl+C to quit")
	}

	return nil
}
