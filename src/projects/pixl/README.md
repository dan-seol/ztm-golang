# PIXL

## About
- Pixel art editor
- Utilize many features of Go
    - Pointers
    - Function Literals / Closures
    - Modules / Packages
    - Interfaces

## Features
- New / Load / Save
- Color picker
- Swatches
- Pan / Zoom images
- Adaptive layout
- Brush indicator

## Structure
- apptype: will have types such as brushes and state
- pixl: main application of the project
- pxcanvas: contain source files for the actual pixel canvas
- swatch: code to display swatches on the screen
- ui: ui
- util: helpers

## PxCanvas
- a single widget
- grey box <- total area consumed by the widget
- x : 0 ->
- y : 0 |
        v
- x offset, y offset

## Swatch

### Wiget in fyne
```
type Widget interface {
    // Base functionality and sate for all widgets
    // (size, position, etc)
    // Initialized with widget.ExtendBaseWidget(widget)
    CanvasObject
    // Creates the renderer which defines how the widget looks
    CreateRenderer() WidgetRenderer
}
```

### Widget Renderer
```
type WidgetRenderer interface {
    // Deprecated
    BackgroundColor() color.Color
    // Internal use: leave empty on implementation
    Destroy()
    // Calculate the position of individual objects
    // that make up this widget
    Layout(size)
    // Minimum dimensions that this widget occupies
    MinSize() Size
    // All objects that should be drawn
    Objects() []CanvasObject
    // An update occurred and the widget needs to be re-drawn
    Refresh()
}
```

### Mouse interfaces
```
type Mouseable interface {
    // Mouse Button pressed
    MouseDown(*MouseEvent)
    // Mouse Button released
    MouseUp(*MouseEvent)
}

type Scrollable interface {
    // Mouse scroll wheel movement
    Scrolled(*ScrollEvent)
}

type Hoverable interface {
// Mouse pointer enters an element. 
    MouseIn(*MouseEvent)
    // Mouse pointer moved over an element
    MouseMoved(*MouseEvent)
    // Mouse pointer leaves an element
    MouseOut()
}
```

### Layout for swatch
- will use gridwrap
install
```
sudo apt install libxcursor-dev
sudo apt install libxrandr-dev
sudo apt-get install libxinerama-dev
sudo apt-get install libxi-dev
sudo apt install libgl-dev
sudo apt -y install pkg-config
``` 