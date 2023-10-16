package main

import (
	"fmt"
	"goexercicio/database"
	"goexercicio/model"
	"time"
)

func main() {
	database.InitDB()
	defer database.Db.Close()

	mercado, data, items := obterDetalhesCompra()
	compra := model.NewCompra(mercado, data, items)

	calcularCustoTotal(&compra)

	fmt.Println("Resumo da compra:")
	fmt.Println(compra.String())
}

func obterDetalhesCompra() (string, time.Time, []model.Item) {
	//Solicitando informações do usuario
	var mercado string
	var data time.Time
	var items []model.Item

	fmt.Print("Nome do mercado: ")
	fmt.Scanln(&mercado)

	fmt.Print("Data da compra (dd/mm/yyyy): ")
	var dataInput string
	fmt.Scanln(&dataInput)
	data, err := time.Parse("02/01/2006", dataInput)
	if err != nil {
		fmt.Println("Erro ao analisar a data:", err)
		return "", time.Time{}, nil
	}

	//Solicitando informaçãoes dos itens
	fmt.Println("Adicione os itens para comprar (digite 'fim' para encerrar)")
	for {
		var nome string
		var quantidade int
		fmt.Print("Nome do item: ")
		fmt.Scanln(&nome)

		if nome == "fim" {
			break
		}

		fmt.Print("Quantidade: ")
		fmt.Scanln(&quantidade)

		items = append(items, model.NewItem(nome, quantidade))
	}

	return mercado, data, items
}

func calcularCustoTotal(compra *model.Compra) {
	query := "SELECT AVG(preco) FROM item_prices WHERE nome = ?"
	stmt, err := database.Db.Prepare(query)
	if err != nil {
		fmt.Printf("Erro ao preparar a consulta: %v\n", err)
		return
	}
	defer stmt.Close()

	//Calculo do custo total
	totalCost := 0.0
	for _, item := range compra.Item {
		
		var media float64
		err :=  stmt.QueryRow(item.Nome).Scan(&media)
		if err != nil {
			fmt.Printf("Erro ao obter a média de preço para %s: %v\n", item.Nome, err)
			return
		}
		itemCost := media * float64(item.Quantidade)
		totalCost += itemCost
	}

	//Adcionando o custo total a estrutura compra
	compra.CustoTotal = totalCost
}