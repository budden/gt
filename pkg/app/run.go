//go:generate gomacro -f -m -w make_fibonacci.gomacro

package app

import (
	"fmt"
)

// Run runs an app
func Run(commandLineArgs []string) {
	fmt.Println("Arg[0] is ")
	fmt.Println(commandLineArgs[0])
}
