package generate

import (
	"gethostdata/options"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	//functions variables
	generate := cli.NewApp()
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "go.dev",
		},
	}
	//cli functions
	generate.Name = "Get Host Data"
	generate.UsageText = ` ./gethostdata <command> --host <your.host>`
	generate.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "IP search",
			Flags:  flags,
			Action: options.SearchIps,
		},
		{
			Name:   "nameserver",
			Usage:  "Nameserver search",
			Flags:  flags,
			Action: options.SearchServers,
		},
		{
			Name:   "txt",
			Usage:  "TXT search",
			Flags:  flags,
			Action: options.SearchTxts,
		},
		{
			Name:   "mx",
			Usage:  "MX search",
			Flags:  flags,
			Action: options.SearchMXs,
		},
	}

	return generate
}
