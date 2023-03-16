package main

import (
	"mini-games/pkg/pendu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Mini-Games!")

	mainContainer := container.NewCenter()
	footer := container.NewHBox(
		widget.NewButton("Pendu", func() {
			mainContainer.RemoveAll()

			entry := widget.NewEntry()
			entry.Resize(fyne.NewSize(100, 50))
			form := &widget.Form{
				Items: []*widget.FormItem{{
					Text:   "Enter a word",
					Widget: entry,
				}},
				OnSubmit: func() {
					mainContainer.RemoveAll()
					mainContainer.Add(pendu.NewPendu(entry.Text))
				},
			}
			mainContainer.Add(
				form,
			)
		}),
	)

	w.SetContent(container.NewVBox(
		mainContainer,
		footer,
	))
	w.Resize(fyne.NewSize(500, 800))

	w.ShowAndRun()
}
