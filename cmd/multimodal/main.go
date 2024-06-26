package main

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
	"github.com/siuyin/aigotut/gfmt"
)

var pathToImage1 = "./dog.png"

func main() {
	cl := client.New()

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

	iter := cl.Model.GenerateContentStream(context.Background(), prompt...)
	gfmt.PrintStreamResponse(iter)
}
