package options

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func SearchIps(host *cli.Context) {
	getHost := host.String("host")
	ips, fail := net.LookupIP(getHost)
	if fail != nil {
		log.Fatal(fail)
	}

	fmt.Printf("\n\nHOST:\n%s\n\nIP's:\n", getHost)

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func SearchServers(host *cli.Context) {
	getHost := host.String("host")
	nameservers, fail := net.LookupNS(getHost)
	if fail != nil {
		log.Fatal(fail)
	}

	fmt.Printf("\n\nHOST:\n%s\n\nNAMESERVER's:\n", getHost)

	for _, ip := range nameservers {
		fmt.Println(ip)
	}
}

func SearchTxts(host *cli.Context) {
	getHost := host.String("host")
	txt, fail := net.LookupTXT(getHost)
	if fail != nil {
		log.Fatal(fail)
	}

	fmt.Printf("\n\nHOST:\n%s\n\nTXT's:\n", getHost)

	for _, ip := range txt {
		fmt.Println(ip)
	}
}

func SearchMXs(host *cli.Context) {
	getHost := host.String("host")
	mx, fail := net.LookupMX(getHost)
	if fail != nil {
		log.Fatal(fail)
	}

	fmt.Printf("\n\nHOST:\n%s\n\nMX's:\n", getHost)

	for _, ip := range mx {
		fmt.Println(ip)
	}
}
