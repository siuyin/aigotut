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
		"Please make it slightly dimmer",
		"Thank you.",
		"Make the room as bright as the day for reading.",
		"Too bright! Make it dimmer.",
		"Much warmer please.",
		"Simulate a cool cloudy day, the light should not be too bright.",
		"Light the room with a warm, very dim, night light.",
	}
	for _, p := range prompts {
		fmt.Printf("prompt: %s\n", p)
		// content := genai.NewUserContent(genai.Text(p))
		// cs.History = append(cs.History, content) // done automatically see: https://github.com/google/generative-ai-go/blob/817706e16e66703730f666180e2577c3738e4bb9/genai/chat.go#L45
		resp, err := cs.SendMessage(context.Background(), genai.Text(p))
		if err != nil {
			log.Fatal(err)
		}

		var (
			fn *genai.FunctionCall
			ok bool
		)
		switch v := resp.Candidates[0].Content.Parts[0].(type) {
		case genai.FunctionCall:
			fn, ok = checkFunctionCall(resp)
			if !ok {
				continue
			}
		case genai.Text:
			gfmt.PrintResponse(resp)
			continue
		default:
			log.Printf("unexpected type: %v", v)
		}

		fnRes := setLightValues(fn.Args["brightness"].(float64), fn.Args["colorTemp"].(string))
		resp, err = cs.SendMessage(context.Background(), fnRes)
		if err != nil {
			log.Fatal(err)
		}

		gfmt.PrintResponse(resp)
		// cs.History = append(cs.History, resp.Candidates[0].Content)
	}

}

func setLightValues(brightness float64, colorTemp string) genai.FunctionResponse {
	fmt.Printf("Setting room brightness to %g at color temperature %q\n---\n", brightness, colorTemp)
	return genai.FunctionResponse{
		Name:     "setLightValues",
		Response: map[string]any{"brightness": brightness, "colorTemp": colorTemp},
	}
}

func checkFunctionCall(resp *genai.GenerateContentResponse) (*genai.FunctionCall, bool) {
	part := resp.Candidates[0].Content.Parts[0]
	funcall, ok := part.(genai.FunctionCall)
	if !ok {
		return nil, false
	}

	if funcall.Name != "setLightValues" {
		log.Printf("ERROR: received bad function call: %s", funcall.Name)
		return nil, false
	}
	return &funcall, true
}
