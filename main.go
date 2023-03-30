package main

import (
	"context"
	"fmt"
	"log"

	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("API_KEY is Missing!")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apikey)

}
