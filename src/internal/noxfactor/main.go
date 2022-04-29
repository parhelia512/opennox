package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/noxworld-dev/opennox/v1/internal/noxfactor/refactor"
)

var (
	fPath = flag.String("path", ".", "path to scan")
)

func main() {
	flag.Parse()
	if err := refactor.Run(*fPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
