package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello Folks to the world of GPT-3")

	godotenv.Load()

	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("API_KEY is Missing!")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apikey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"one tweet on golang is: "},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Choices[0].Text)

}
