package main

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	consoleLog = zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "", NoColor: true})
)

func main() {
	
}
