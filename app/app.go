package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

func Generator() *cli.App {
	app := cli.NewApp()

	app.Name = "Plicação de linha de comando"
	app.Usage = "Busca ips e nomes de servidores webs"
	app.Commands = []*cli.Command{
		{
			Name:  "ip",
			Usage: "Busca de ips de indereços na internet",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "host",
					Value: "google.com",
					Usage: "Host a ser pesquisado",
				},
			},
			Action: searchIps,
		},
	}
	return app
}

func searchIps(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}
