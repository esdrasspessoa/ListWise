package model

import (
	"fmt"
	"time"
)

type Compra struct {
	Data       time.Time
	Mercado    string
	Item       []Item
	CustoTotal float64
}

func NewCompra(mercado string, data time.Time, itens []Item) Compra {
	return Compra{
		Data:    data,
		Mercado: mercado,
		Item:    itens,
	}
}

func (c Compra) String() string {
    return fmt.Sprintf("Data da Compra: %s\nMercado: %s\nItens para comprar:\n%s\nCusto Total: %.2f",
        c.Data.Format("02/01/2006"), c.Mercado, c.itemsString(), c.CustoTotal)
}

func (c Compra) itemsString() string {
	var itemsStr string
	for _, item := range c.Item {
		itemsStr += fmt.Sprintf("- %s\n", item)
	}
	return itemsStr
}
