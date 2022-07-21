package generate

import (
	"gethostdata/options"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	generate := cli.NewApp()
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "go.dev",
		},
	}
	generate.Name = "Get Host Data"
	generate.UsageText = ` ./gethostdata <command> --host <your.host>`
	generate.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "IP search",
			Flags:  flags,
			Action: options.GetInformation,
		},
		{
			Name:   "nameserver",
			Usage:  "Nameserver search",
			Flags:  flags,
			Action: options.GetInformation,
		},
		{
			Name:   "txt",
			Usage:  "TXT search",
			Flags:  flags,
			Action: options.GetInformation,
		},
		{
			Name:   "mx",
			Usage:  "MX search",
			Flags:  flags,
			Action: options.GetInformation,
		},
	}

	return generate
}
