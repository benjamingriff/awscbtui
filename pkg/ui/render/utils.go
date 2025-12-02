package render

import (
	"github.com/awesome-gocui/gocui"
)

func ensureVisibleWithHeader(v *gocui.View, idx, n, headerLines int) (int, error) {
  if n <= 0 {
    _ = v.SetOrigin(0, 0)
    _ = v.SetCursor(0, headerLines) // place cursor on first data row
    return 0, nil
  }
  if idx < 0 { idx = 0 }
  if idx >= n { idx = n - 1 }

  ox, oy := v.Origin()
  _, h := v.Size()

  // Translate idx (data row) to absolute view row
  absRow := headerLines + idx

  // Visible window in absolute rows is [oy, oy+h)
  if absRow >= oy && absRow < oy+h {
    cy := absRow - oy
    if cy < 0 { cy = 0 }
    if cy >= h { cy = h - 1 }
    _ = v.SetCursor(0, cy)
    return idx, nil
  }

  // Scroll up so absRow is first visible data row under header
  if absRow < oy {
    // Try to place header at top; ensure data row appears
    newOy := absRow - headerLines
    if newOy < 0 { newOy = 0 }
    if err := v.SetOrigin(ox, newOy); err != nil { return idx, err }
    _ = v.SetCursor(0, absRow-newOy)
    return idx, nil
  }

  // Scroll down so absRow is the last visible row
  newOy := absRow - (h - 1)
  if newOy < 0 { newOy = 0 }
  if err := v.SetOrigin(ox, newOy); err != nil { return idx, err }
  _ = v.SetCursor(0, absRow-newOy)
  return idx, nil
}
