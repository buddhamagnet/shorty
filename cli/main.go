package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buddhamagnet/shorty/shorty"
	"github.com/joho/godotenv"
	flag "github.com/spf13/pflag"
)

var longURL, service, id string

func init() {
	flag.StringVarP(&longURL, "url", "u", "", "URL to shorten")
	flag.StringVar(&id, "id", "", "shortened URL ID to decode")
	flag.StringVarP(&service, "service", "s", "shorten", "service to invoke")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()

	switch service {
	case "shorten":
		if longURL == "" {
			fmt.Println("usage: cli (-u|--url)=<url>")
			os.Exit(1)
		}
		id, err := shorty.Shorten(longURL)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("shortened URL: http://localhost:8080/" + id + "\n")

	case "decode":
		if id == "" {
			fmt.Println("usage: cli (--id)=<id>")
			os.Exit(1)
		}
		longURL, err := shorty.Decode(id)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("long URL: " + longURL + "\n")
	}

}
