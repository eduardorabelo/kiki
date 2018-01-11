package main

import (
	"flag"
	"fmt"

	"github.com/schollz/kiki/src/logging"
)

func main() {
	port := flag.String("port", "8003", "port for the data (this) server")
	debug := flag.Bool("debug", false, "turn on debug mode")
	location := flag.String("path", ".", "path to the kiki database folder")
	flag.Parse()
	if *debug {
		logging.Log.Debug(true)
	} else {
		logging.Log.Debug(false)
	}
	Port = *port
	Location = *location
	err := Run()
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}
