package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
}

func ex1() {
	log.Debug().Bool("bool val", true).Time("time val", time.Now()).Msg("debug msg")
	log.Info().Bool("bool val", true).Time("time val", time.Now()).Send()
}

func ex2() {
	logger := zerolog.New(os.Stdout).With().Time("time val", time.Now()).Logger()
	logger.Info().Send()
}

func ex3() {
	makeErr := func() error { return fmt.Errorf("some error: %w", errors.New("sub-error")) }

	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return errors.Unwrap(err)
	}

	log.Err(makeErr()).Send()

	logger := log.With().Stack().Logger()
	logger.Err(makeErr()).Send()
}

func ex4() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	logger.Info().Bool("bool val", true).Msg("info msg")
}
