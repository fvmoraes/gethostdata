package options

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GetInformation(data *cli.Context) {
	getHost := data.String("host")
	getCommand := data.Command.Name

	fmt.Println(getHost, getCommand)

	switch getCommand {
	case "ip":
		informations, fail := net.LookupIP(getHost)
		testFail(fail)
		dataHeader(getHost)
		dataSearch(informations)
	case "nameserver":
		informations, fail := net.LookupNS(getHost)
		testFail(fail)
		dataHeader(getHost)
		dataSearch(informations)
	case "txt":
		informations, fail := net.LookupTXT(getHost)
		testFail(fail)
		dataHeader(getHost)
		dataSearch(informations)
	case "mx":
		informations, fail := net.LookupMX(getHost)
		testFail(fail)
		dataHeader(getHost)
		dataSearch(informations)
	}
}

func testFail(fail error) {
	if fail != nil {
		log.Fatal(fail)
	}
}

func dataHeader(host string) {
	fmt.Printf("\n\nHOST:\n%s\n\nINFORMATION:\n", host)
}

func dataSearch[GenericType any](informations []GenericType) {
	for _, information := range informations {
		fmt.Println(information)
	}
}
