package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/Gyabi/go-translateApp/translatePost"
)

type MyMainWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
	results  *walk.TextEdit
}

func (mw *MyMainWindow) clickedUpArrow() {
	text := mw.results.Text()

	output := translatePost.Translate_post(text, "ja", "en")

	mw.textArea.SetText(output)
}

func (mw *MyMainWindow) clickedDownArrow() {
	text := mw.textArea.Text()

	output := translatePost.Translate_post(text, "en", "ja")

	mw.results.SetText(output)
}

func main() {
	mw := &MyMainWindow{}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "todo",
		MinSize:  Size{300, 400},
		Layout:   VBox{},
		Children: []Widget{
			// GroupBox{
			// 	Layout: HBox{},
			// 	Children: []Widget{
			// 		Label{
			// 			Text: "Please enter the text you want to translate",
			// 		},
			// 		PushButton{
			// 			Text:      "translate",
			// 			OnClicked: mw.clicked,
			// 		},
			// 	},
			// },
			Label{
				Text: "Please enter the text you want to translate",
			},
			GroupBox{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "English",
					},
					TextEdit{
						AssignTo: &mw.textArea,
					},
				},
			},
			HSplitter{
				StretchFactor: 20,
				Children: []Widget{
					PushButton{
						Text:      "↓",
						OnClicked: mw.clickedDownArrow,
					},
					HSpacer{},
					PushButton{
						Text:      "↑",
						OnClicked: mw.clickedUpArrow,
					},
				},
			},
			GroupBox{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "Japanese",
					},
					TextEdit{
						AssignTo: &mw.results,
					},
				},
			},
			// TextEdit{
			// 	AssignTo: &mw.textArea,
			// },

			// TextEdit{
			// 	AssignTo: &mw.results,
			// 	Row:      5,
			// },
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
