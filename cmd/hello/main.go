package main

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
	"github.com/siuyin/aigotut/gfmt"
)

func main() {
	cl := client.New()
	defer cl.Close()

	resp, err := cl.Model.GenerateContent(context.Background(), genai.Text("What is the significance of the number 42?"))
	if err != nil {
		log.Fatal(err)
	}

	gfmt.PrintResponse(resp)
}
