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

	// start chat session
	cs := cl.Model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Hello, I have 2 dogs in my house."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Great to meet you. What would you like to know?"),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(context.Background(), genai.Text("How many paws are in my house?"))
	if err != nil {
		log.Fatal(err)
	}

	gfmt.PrintResponse(resp)
}
