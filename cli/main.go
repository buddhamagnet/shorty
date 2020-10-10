package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/buddhamagnet/shorty/cli/flags"
	"github.com/buddhamagnet/shorty/shorty"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, output, err := flags.Parse(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	switch cfg.Service {
	case "shorten":
		id, err := shorty.Shorten(cfg.LongURL)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("shortened URL: http://localhost:8080/" + id + "\n")
	case "decode":
		longURL, err := shorty.Decode(cfg.ID)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Printf("long URL: " + longURL + "\n")
	}

}
