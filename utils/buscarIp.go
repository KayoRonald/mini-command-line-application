package utils

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func BuscarIp(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
