package main

import (
	"fmt"
	app "linha-de-comando/App"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
