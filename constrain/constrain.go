package constrain

import (
	"os"

	"github.com/rs/zerolog"
)

var ConsoleLog = zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "", NoColor: true})
