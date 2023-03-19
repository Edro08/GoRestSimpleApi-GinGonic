package main

import (
	"GoRestSimpleApi/cmd/bootstrap"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
