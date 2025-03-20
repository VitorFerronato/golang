package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "vitorferronato:123456@/golang?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão esta abertaq!")

}
