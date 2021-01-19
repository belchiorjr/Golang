package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"


	flags := [] cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}


	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca de IPS de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

/* Faz a Busca */
func buscarIps(c *cli.Context) {
	host := c.String("host")

	// Pacote net
	ips, erro := net.LookupIP(host)

	if erro != nil{
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

/* Faz a do servidor */
func buscarServidores (c *cli.Context) {
	host := c.String("host")

	// Pacote net
	servidores, erro := net.LookupNS(host)

	if erro != nil{
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}