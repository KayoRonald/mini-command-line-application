package app

import (
	"command-line/utils"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscar IPs e nomes de servidores"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "buscar endereços IPS na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "golang.org",
				},
			},
			Action: utils.BuscarIp,
		},
	}
	return app
}
