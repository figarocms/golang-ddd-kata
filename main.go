package main

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/bootstrap"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Get Things Done Todos Rest API
// @version 1.0
// @description This is a golang project that serves as an example of hexagonal architecture

// @contact.email gsigety@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
func main() {
	bootstrap.LaunchApp()
}

func init() {
	// Log as JSON instead of the default ASCII formatter.

	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}
