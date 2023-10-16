package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:191103@tcp(127.0.0.1:3306)/itens_db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Erro ao verificar conexão com o banco de dados:", err)
		return
	}
}
