package pendu

import (
	"fyne.io/fyne/v2"
)

type Test struct {
}

func NewTest(res fyne.Resource) *Test {
	essai := &Test{}
	essai.ExtendBaseWidget(essai)
	essai.SetResource(res)

	return essai
}
