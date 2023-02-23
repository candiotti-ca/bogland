package persistance

import (
	"fmt"
	"youpi/pkg/kida"
)

type InMemory struct {
}

func (InMemory) Save(k kida.Kida) error {
	fmt.Println("save kida:", k.Id)
	return nil
}

func (InMemory) Retrieve(id string) (kida.Kida, error) {
	return kida.Kida{}, nil
}
