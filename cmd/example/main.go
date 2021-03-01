package main

import (
	"strings"

	"github.com/lxn/walk"
	decl "github.com/lxn/walk/declarative"
)

// Note: Instead of running this with go run, compile with Makefile.
func main() {
	var inTE, outTE *walk.TextEdit

	decl.MainWindow{
		Title: "Capitalizer",
		MinSize: decl.Size{
			Width:  320,
			Height: 240,
		},
		Layout: decl.VBox{},
		Children: []decl.Widget{
			decl.HSplitter{
				Children: []decl.Widget{
					decl.TextEdit{AssignTo: &inTE},
					decl.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			decl.PushButton{
				Text: "Capitalize",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
