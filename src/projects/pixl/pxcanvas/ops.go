package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff

	pxCanvas.Refresh()
}

func (PxCanvas *PxCanvas) Scale(direction int) {
	switch {
	case direction > 0:
		PxCanvas.PxSize++
	case direction < 0:
		if PxCanvas.PxSize > 2 {
			PxCanvas.PxSize--
		}
	default:
		PxCanvas.PxSize = 10
	}
}
