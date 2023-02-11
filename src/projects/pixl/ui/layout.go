package ui

import (
	"fyne.io/fyne/v2/container"
)

func Setup(app *AppInit) {
	// phase 4 - menus
	SetupMenus(app)
	// phase 1 - swatches
	swatchesContainer := BuildSwatches(app)
	// phase 2 - color picker
	colorPicker := SetUpColorPicker(app)
	// top, bottom, left, right
	appLayout := container.NewBorder(
		nil, swatchesContainer, nil, colorPicker, app.PixlCanvas)

	//app.PixlWindow.SetContent(swatchesContainer)
	app.PixlWindow.SetContent(appLayout)
}
