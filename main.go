package main

import (
	"gethostdata/generate"
	"log"
	"os"
)

func main() {
	application := generate.Generate()
	if fail := application.Run(os.Args); fail != nil {
		log.Fatal(fail)
	}
}
