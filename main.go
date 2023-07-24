package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
)

const version = "v0.3"

const usage = `A super-simple file server.

Usage:
  filesrvr -r ROOTDIR -p PORT [-a USER:PASS]
  filesrvr -h | --help
  filesrvr --version

Options:
  -h, --help              Show this screen.
  --version               Show version.
  -r, --rootdir DIR       Root directory for files.
  -p, --port PORT         Port to listen on.
  -a, --auth USER:PASS    Enable basic auth protection, using user/pass combo.

Homepage: https://github.com/alasdairmorris/filesrvr
`

func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type application struct {
	config struct {
		rootdir  string
		port     int
		username string
		password string
	}
}

func main() {

	app := new(application)

	app.init()
	app.run()

}

func (app *application) init() {

	var (
		opts       docopt.Opts
		rootdir    string
		absRootDir string
		port       string
		auth       string
		err        error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	exitOnError(err)

	// Rootdir
	rootdir, err = opts.String("--rootdir")
	exitOnError(err)

	absRootDir, err = filepath.Abs(rootdir)
	exitOnError(err)

	app.config.rootdir = absRootDir

	// Port
	port, err = opts.String("--port")
	exitOnError(err)

	app.config.port, err = strconv.Atoi(port)
	exitOnError(err)

	// Auth
	auth, err = opts.String("--auth")

	if len(auth) > 0 {
		bits := strings.Split(auth, ":")
		if len(bits) == 2 {
			app.config.username = bits[0]
			app.config.password = bits[1]
		} else {
			fmt.Fprintf(os.Stderr, "Error parsing auth details %q\n", auth)
			os.Exit(1)
		}
	}
}

func (app *application) basicAuth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.config.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.config.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				http.FileServer(http.Dir(app.config.rootdir)).ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (app *application) run() {
	log.Printf("Starting server on port %d, with rootdir %s", app.config.port, app.config.rootdir)

	if app.config.username != "" && app.config.password != "" {
		http.Handle("/", app.basicAuth())
	} else {
		http.Handle("/", http.FileServer(http.Dir(app.config.rootdir)))
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(app.config.port), nil))
}
