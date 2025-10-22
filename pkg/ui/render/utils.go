package render

import (
	"github.com/awesome-gocui/gocui"
)

func ensureVisible(v *gocui.View, idx int, n int) (int, error) {
	if n <= 0 {
		_ = v.SetOrigin(0, 0)
		_ = v.SetCursor(0, 0)
		return 0, nil
	}
	if idx < 0 {
		idx = 0
	}
	if idx >= n {
		idx = n - 1
	}

	ox, oy := v.Origin()
	_, h := v.Size()

	// If already visible, keep current origin
	if idx >= oy && idx < oy+h {
		cy := idx - oy
		if cy < 0 {
			cy = 0
		}
		if cy >= h {
			cy = h - 1
		}
		_ = v.SetCursor(0, cy)
		return idx, nil
	}

	// If above current window, scroll up
	if idx < oy {
		newOy := idx
		if newOy < 0 {
			newOy = 0
		}
		if err := v.SetOrigin(ox, newOy); err != nil {
			return idx, err
		}
		_ = v.SetCursor(0, 0)
		return idx, nil
	}

	// Else below current window: scroll so idx is last visible row
	newOy := idx - (h - 1)
	if newOy < 0 {
		newOy = 0
	}
	if err := v.SetOrigin(ox, newOy); err != nil {
		return idx, err
	}
	_ = v.SetCursor(0, idx-newOy)
	return idx, nil
}

