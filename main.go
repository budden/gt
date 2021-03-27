package main

import (
	"fmt"
	"os"

	"github.com/budden/gt/pkg/app"
)

func main() {
	fmt.Println("fibint32(8) = ", app.Fibint32(8))
	app.Run(os.Args)
}
