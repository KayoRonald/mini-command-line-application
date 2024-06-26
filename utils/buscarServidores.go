package utils

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func BuscarNomeServidor(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
