package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
)

var lightControlTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{{
		Name:        "setLightValues",
		Description: "Set the brightness and color temperature of a room light.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"brightness": {
					Type:        genai.TypeString,
					Description: "Light level from 0 to 100. Zero is off and 100 is full brightness.",
				},
				"colorTemp": {
					Type:        genai.TypeString,
					Description: "Color Temperature of the light fixture which can be daylight, cool or warm.",
				},
			},
			Required: []string{"brightness"},
		},
	}},
}

func main() {
	cl := client.New()
	defer cl.Close()

	cl.Model.Tools = []*genai.Tool{lightControlTool}

	// Start chat session.
	cs := cl.Model.StartChat()
	prompt := "Dim the light so that the room feels cosy and warm."
	//prompt := "Make the room as bright as the day for reading."
	//prompt := "Simulate a cool cloudy day, the light should not be too bright."
	//prompt := "Light the room with a warm, very dim, night light."
	resp, err := cs.SendMessage(context.Background(), genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	// check response include function call
	part := resp.Candidates[0].Content.Parts[0]
	funcall, ok := part.(genai.FunctionCall)
	if !ok {
		log.Fatalf("Expected type FunctionCall, got %T", part)
	}
	if g, e := funcall.Name, lightControlTool.FunctionDeclarations[0].Name; g != e {
		log.Fatalf("Expected FunctionCall.Name %q, got %q", e, g)
	}
	fmt.Printf("Received function call response:\n%q\n\n", part)

}

func setLightValues(brightness int, colorTemp string) map[string]any {
	return map[string]any{
		"brightness": brightness,
		"colorTemp":  colorTemp}
}
