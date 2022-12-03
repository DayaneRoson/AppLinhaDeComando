package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca por IPs e Nomes de Servidor na Internet"
	app.Author = "Day"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca por IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "bibliaonline.com.br/acf",
				},
			},
			Action: buscarIps,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {

	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
