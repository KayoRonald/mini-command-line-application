package app

import (
	"command-line/utils"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscar IPs e nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "golang.org",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca endereços IPS na internet",
			Flags: flags,
			Action: utils.BuscarIp,
		},
		{
			Name:  "servidores",
			Usage: "Busca o nome do servidores na internet",
			Flags: flags,
			Action: utils.BuscarNomeServidor,
		},
	}
	return app
}
