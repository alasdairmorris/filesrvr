package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `A super-simple file server.

Usage:
  filesrvr
  filesrvr -h | --help
  filesrvr --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/filesrvr v0.0.1")
	fmt.Println(opts)
}
