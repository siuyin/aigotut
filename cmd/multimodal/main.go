package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var pathToImage1 = "./dog.png"

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts
	model := client.GenerativeModel("gemini-1.5-flash")

	imgData1, err := os.ReadFile(pathToImage1)
	if err != nil {
		log.Fatal(err)
	}

	// imgData2, err := os.ReadFile(pathToImage1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	prompt := []genai.Part{
		genai.ImageData("png", imgData1),
		// genai.ImageData("jpeg", imgData2),
		genai.Text("Describe in detail the contents of the image."),
	}

	iter := model.GenerateContentStream(ctx, prompt...)

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
