package render

import (
	"fmt"
	"time"
	"strconv"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderBuilds(v *gocui.View, s *state.AppState) {
	v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	v.Title = "[2]-Builds"
	v.Wrap = false
	v.Clear()

	v.FrameColor = gocui.ColorDefault
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen

	fmt.Fprintln(v, " ID                                                       | BUILD NUMBER | STATUS       | DURATION | ")
  fmt.Fprintln(v, " -------------------------------------------------------- | ------------ | ------------ | -------- | ")

	v.Highlight = (s.UI.FocusedView == state.ViewBuilds)

	builds := s.Data.Builds[s.UI.SelectedProjectName]
		for _, b := range builds {
			id := right(b.ID, 56)
			bn := strconv.FormatInt(b.BuildNumber, 12)
			dur := humanDuration(buildDuration(b.StartTime, b.EndTime, b.Status))
			fmt.Fprintf(v, " %-56s   %-12s   %-12s   %8s\n", id, bn, colorStatus(b.Status), dur)
		}

	if v.Highlight {
		v.FrameColor = gocui.ColorGreen
		idx := s.UI.FocusedBuildIdx
		_, _ = ensureVisibleWithHeader(v, idx, len(builds), 2)
	} else {
		_ = v.SetCursor(0, 0)
	}
}

func colorStatus(s string) string {
  switch s {
  case "SUCCEEDED":
    return "\x1b[32mSUCCEEDED\x1b[0m"
  case "FAILED", "FAULT", "TIMED_OUT", "STOPPED":
    return "\x1b[31m" + s + "\x1b[0m"
  case "IN_PROGRESS":
    return "\x1b[33mIN_PROGRESS\x1b[0m"
  default:
    return s
  }
}

func buildDuration(start, end time.Time, status string) time.Duration {
  if start.IsZero() {
    return 0
  }
  if end.IsZero() && (status == "IN_PROGRESS") {
    return time.Since(start)
  }
  if end.After(start) {
    return end.Sub(start)
  }
  return 0
}

func humanDuration(d time.Duration) string {
  if d <= 0 {
    return "-"
  }
  if d < time.Minute {
    return fmt.Sprintf("%ds", int(d.Seconds()))
  }
  if d < time.Hour {
    m := int(d.Minutes())
    s := int(d.Seconds()) % 60
    return fmt.Sprintf("%dm%02ds", m, s)
  }
  h := int(d.Hours())
  m := int(d.Minutes()) % 60
  return fmt.Sprintf("%dh%02dm", h, m)
}

func right(s string, width int) string {
  if len(s) <= width {
    return s
  }
  if width <= 1 {
    return "…"
  }
  return "…" + s[len(s)-width+1:]
}

func setColumnWidth(columnName string, width int) int {
	nameSize := len(columnName)
	return nameSize
}
	


