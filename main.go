package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/docopt/docopt-go"
)

const version = "v0.1"

const usage = `A super-simple file server.

Usage:
  filesrvr -r ROOTDIR -p PORT
  filesrvr -h | --help
  filesrvr --version

Options:
  -h, --help              Show this screen.
  --version               Show version.
  -r, --rootdir DIR       Root directory for files.
  -p, --port PORT         Port to listen on.

Homepage: https://github.com/alasdairmorris/filesrvr
`

type Config struct {
	Rootdir string
	Port    int
}

var config Config

func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// Parse and validate command-line arguments
func getConfig() Config {

	var (
		retval     Config
		opts       docopt.Opts
		rootdir    string
		absRootDir string
		port       string
		err        error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	exitOnError(err)

	// Rootdir
	rootdir, err = opts.String("--rootdir")
	exitOnError(err)

	absRootDir, err = filepath.Abs(rootdir)
	exitOnError(err)

	retval.Rootdir = absRootDir

	// Port
	port, err = opts.String("--port")
	exitOnError(err)

	retval.Port, err = strconv.Atoi(port)
	exitOnError(err)

	return retval
}

func main() {

	config = getConfig()

	log.Printf("Starting server on port %d, with rootdir %s", config.Port, config.Rootdir)

	http.Handle("/", http.FileServer(http.Dir(config.Rootdir)))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
}
