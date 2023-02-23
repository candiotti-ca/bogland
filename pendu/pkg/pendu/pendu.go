package pendu

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Pendu struct {
	word  string
	lives int
	tries []string
	found []string
}

func NewPendu(word string) fyne.CanvasObject {
	pendu := Pendu{word: word, lives: 5, tries: []string{}, found: []string{}}
	penduContainer := container.NewVBox()
	letters := container.NewHBox()
	for i := 0; i < len(word); i++ {
		letters.Add(canvas.NewLine(color.White))
	}

	entry := widget.NewEntry()
	entry.OnChanged = func(s string) {
		if len(s) > 1 {
			entry.Text = string(s[1])
			entry.Refresh()
		}
		pendu.Try(entry.Text)
		pendu.Refresh(letters)
		game := pendu.Verify()

		if game > 0 { //win

		} else if game < 0 { //loose

		}
	}
	form := widget.NewForm(widget.NewFormItem("Enter a letter", entry))
	penduContainer.Add(letters)
	penduContainer.Add(form)

	return penduContainer
}

func (p *Pendu) Try(char string) {
	if strings.Contains(p.word, char) {
		p.found = append(p.found, char)
	} else if !Contains(p.tries, char) {
		p.tries = append(p.tries, char)
		p.lives--
	}
}

func (p Pendu) Refresh(letters *fyne.Container) {
	letters.RemoveAll()
	for i := 0; i < len(p.word); i++ {
		letter := string(p.word[i])
		if Contains(p.found, letter) {
			letters.Add(widget.NewLabel(letter))
		} else {
			letters.Add(canvas.NewLine(color.White))
		}
	}
}

func Contains(slice []string, element string) bool {
	result := false
	for i := 0; i < len(slice); i++ {
		result = result || strings.EqualFold(string(slice[i]), element)
	}
	return result
}

func (p Pendu) Verify() int {
	if p.lives < 1 {
		return -1
	}

	allMatch := true
	for i := 0; i < len(p.word); i++ {
		allMatch = allMatch && Contains(p.found, string(p.word[i]))
	}

	if allMatch {
		return 1
	}

	return 0
}
