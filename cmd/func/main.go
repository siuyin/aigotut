package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
	"github.com/siuyin/aigotut/gfmt"
)

var lightControlTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{{
		Name:        "setLightValues",
		Description: "Set the brightness and color temperature of a room light.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"brightness": {
					Type:        genai.TypeInteger,
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
	prompts := []string{
		"Dim the light so that the room feels cosy and warm. Let me know when it is done",
		"Make the room as bright as the day for reading.",
		"Simulate a cool cloudy day, the light should not be too bright.",
		"Light the room with a warm, very dim, night light.",
	}
	for _, p := range prompts {
		fmt.Printf("prompt: %s\n", p)
		resp, err := cs.SendMessage(context.Background(), genai.Text(p))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("number of candidates in response: %d\n", len(resp.Candidates))
		// check response include function call
		part := resp.Candidates[0].Content.Parts[0]
		funcall, ok := part.(genai.FunctionCall)
		if !ok {
			log.Printf("Expected type FunctionCall, got %T", part)
			continue
		}

		if funcall.Name != "setLightValues" {
			log.Printf("ERROR: received bad function call: %s", funcall.Name)
			continue
		}

		brightness, colorTemp:=setLightValues(funcall.Args["brightness"].(float64), funcall.Args["colorTemp"].(string))
		resp,err=cs.SendMessage(context.Background(),genai.FunctionResponse{
			Name: funcall.Name,
			Response: map[string]any{"brightness":brightness,"colorTemp":colorTemp},
		})
		if err != nil {
			log.Fatal(err)
		}
	
		gfmt.PrintResponse(resp)
	}

}

func setLightValues(brightness float64, colorTemp string) (float64, string) {
	fmt.Printf("Setting room brightness to %g at color temperature %q\n---\n", brightness, colorTemp)
	return brightness, colorTemp
}
