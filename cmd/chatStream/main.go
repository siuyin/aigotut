package main

import (
	"context"
	"fmt"
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

	interactiveSession(cs)

}

func interactiveSession(cs *genai.ChatSession) {
	iter := cs.SendMessageStream(context.Background(), genai.Text("How many paws are in my house?"))
	gfmt.PrintStreamResponse(iter)
	cs.History = append(cs.History, iter.MergedResponse().Candidates[0].Content)

	prompts := []string{
		"There are 4 people in the house. How many feet and paws are there now?",
		"Now add two cats",
	}
	for _, p := range prompts {
		fmt.Println(p)
		res, err := cs.SendMessage(context.Background(), genai.Text(p))
		if err != nil {
			log.Fatal(err)
		}
		gfmt.PrintResponse(res)
		cs.History = append(cs.History, res.Candidates[0].Content)
	}
}
