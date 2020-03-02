package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	f := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		port     = f.String("p", "", "Port this application will run on")
		proxyURL = f.String("u", "", "Proxy URL")
		sleep    = f.Duration("s", 0, "Duration to sleep for (ex: 50us, 1000ms, 1s, 1h)")
	)

	if err := f.Parse(args[1:]); err != nil {
		return err
	}

	if *port == "" {
		*port = os.Getenv("PORT")
		if *port == "" {
			return errors.New("missing port")
		}
	}

	if *proxyURL == "" {
		*proxyURL = os.Getenv("PROXY_URL")
		if *proxyURL == "" {
			return errors.New("missing proxy url")
		}
	}

	if *sleep == 0 {
		s, err := time.ParseDuration(os.Getenv("SLEEP_DURATION"))
		if err != nil || s == 0 {
			return errors.New("missing sleep duration")
		}
		*sleep = s
	}

	if err := startServer(*port, *proxyURL, *sleep); err != nil {
		return err
	}

	return nil
}
