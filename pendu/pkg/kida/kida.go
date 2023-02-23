package kida

import (
	"strconv"

	"fyne.io/fyne/v2/data/binding"
)

type SaverKida interface {
	Save(k Kida) error
}

type Kida struct {
	Id          string
	count       int
	PetCount    binding.String
	persistance SaverKida
}

func NewKida(id string, saver SaverKida) Kida {
	k := Kida{
		Id:          id,
		PetCount:    binding.NewString(),
		persistance: saver,
	}
	k.PetCount.Set("0")
	return k
}

func (k *Kida) Pet() {
	k.count++
	k.PetCount.Set(strconv.Itoa(k.count))
	k.persistance.Save(*k)
}
