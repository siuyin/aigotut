package main

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
	"github.com/siuyin/aigotut/gfmt"
)

func main() {
	cl := client.New()
	defer cl.Close()

	iter := cl.Model.GenerateContentStream(context.Background(), genai.Text("What is the significance of the number 42?"))

	gfmt.PrintStreamResponse(iter)
}
