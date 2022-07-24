package options

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GetInformation(data *cli.Context) {
	hostName := data.String("host")
	commandFlag := data.Command.Name

	switch commandFlag {
	case "ip":
		informations, fail := net.LookupIP(hostName)
		getHostInfo(hostName, informations, fail)
	case "nameserver":
		informations, fail := net.LookupNS(hostName)
		getHostInfo(hostName, informations, fail)
	case "txt":
		informations, fail := net.LookupTXT(hostName)
		getHostInfo(hostName, informations, fail)
	case "mx":
		informations, fail := net.LookupMX(hostName)
		getHostInfo(hostName, informations, fail)
	}
}

func getHostInfo[GenericType any](host string, informations []GenericType, fail error) {
	fmt.Printf("\n\nHOST:\n%s\n\nINFORMATION:\n", host)

	for _, information := range informations {
		fmt.Println(information)
	}

	if fail != nil {
		log.Fatal(fail)
	}
}
