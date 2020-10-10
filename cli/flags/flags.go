package flags

import (
	"bytes"
	"errors"

	flag "github.com/spf13/pflag"
)

// Config represents the combined configuration data for the CLI.
type Config struct {
	LongURL string
	Service string
	ID      string
	args    []string
}

// Parse handles all flags and their combinations.
func Parse(program string, args []string) (config *Config, output string, err error) {
	flags := flag.NewFlagSet(program, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var cfg Config

	flags.StringVarP(&cfg.LongURL, "url", "u", "", "URL to shorten")
	flags.StringVar(&cfg.ID, "id", "", "shortened URL ID to decode")
	flags.StringVarP(&cfg.Service, "service", "s", "shorten", "service to invoke")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	switch cfg.Service {
	case "shorten":
		if cfg.LongURL == "" {
			return nil, buf.String(), errors.New("usage: cli (--url|-u)=<url>")
		}

	case "decode":
		if cfg.ID == "" {
			if cfg.ID == "" {
				return nil, buf.String(), errors.New("usage: cli (--id)=<id>")
			}
		}
	}
	cfg.args = flags.Args()
	return &cfg, buf.String(), nil
}
