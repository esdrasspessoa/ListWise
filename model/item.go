package model

import "fmt"

type Item struct {
	Nome       string
	Quantidade int
}

func NewItem(nome string, quantidade int) Item {
	return Item{
		Nome:       nome,
		Quantidade: quantidade,
	}
}

func (i Item) String() string {
	return fmt.Sprintf("%s (Quantidade: %d)", i.Nome, i.Quantidade)
}