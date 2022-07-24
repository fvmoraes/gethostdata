package options

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

var hostName string
var commandFlag string

func GetHostInfo(data *cli.Context) {
	hostName = data.String("host")
	commandFlag = data.Command.Name

	switch commandFlag {
	case "ip":
		printHostInfo(net.LookupIP(hostName))
	case "nameserver":
		printHostInfo(net.LookupNS(hostName))
	case "txt":
		printHostInfo(net.LookupTXT(hostName))
	case "mx":
		printHostInfo(net.LookupMX(hostName))
	}
}

func printHostInfo[GenericType any](information []GenericType, fail error) {
	fmt.Printf("\n\nHOST:\n%s\n\nINFORMATION:\n", hostName)

	for _, it := range information {
		fmt.Println(it)
	}

	if fail != nil {
		log.Fatal(fail)
	}
}
