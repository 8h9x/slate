package server

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type Proxy struct {
}

func NewProxy() *Proxy {
    return &Proxy{}
}

func (b *Proxy) Widget() fyne.CanvasObject {
    return container.NewVBox(
        widget.NewLabel("Running on port :5555"),
        widget.NewButton("Start Proxy", nil),
    )
}
