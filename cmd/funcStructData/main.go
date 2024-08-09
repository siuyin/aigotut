package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
)

type entities struct {
	People        []person `json:"people"`
	Places        []place
	Things        []thing
	Relationships []relationship
}

type person struct {
	Name           string
	Description    string
	StartPlaceName string `json:"start_place_name"`
	EndPlaceName   string `json:"end_place_name"`
}

type place struct {
	Name        string
	Description string
}

type thing struct {
	Name           string
	Description    string
	StartPlaceName string `json:"start_place_name"`
	EndPlaceName   string `json:"end_place_name"`
}

type relationship struct {
	Person1Name  string `json:"person_1_name"`
	Person2Name  string `json:"person_2_name"`
	Relationship string
}

var (
	story string
	cl    *client.Info
	ctx   context.Context
)

func init() {
	cl = client.New()
	ctx = context.Background()
	genStory()
}

func main() {
	// fmt.Println(story)
	str := extractEntities()
	var ent entities
	if err := json.Unmarshal([]byte(str), &ent); err != nil {
		log.Printf("Unmarshal entities: %v\n%s", err, str)
		return
	}
	fmt.Printf("%#v\n", ent)
}

func genStory() {
	b, err := os.ReadFile("story.txt")
	if err == nil {
		story = string(b)
		return
	}

	resp, err := cl.Model.GenerateContent(ctx, genai.Text(`Write a long story
	about a girl with a magic backback, her family and at least
	on other character. Make sure every person has Chinese names and
	every place has names.
	Don't forget to describe the contents of the backpack,
	and where everyone and everything starts and ends up.`))
	if err != nil {
		log.Printf("GenerateContent: %v", err)
	}

	story = string(resp.Candidates[0].Content.Parts[0].(genai.Text))
	if err := os.WriteFile("story.txt", []byte(story), 0640); err != nil {
		log.Printf("could not create story.txt: %v", err)
	}
}

func extractEntities() string {
	resp, err := cl.Model.GenerateContent(ctx, genai.Text(fmt.Sprintf(
		`Please return JSON describing
		the people, places, things and relationships from this story using the following schema:

	{"people": list[PERSON], "places":list[PLACE], "things":list[THING], "relationships": list[RELATIONSHIP]}
    PERSON = {"name": str, "description": str, "start_place_name": str, "end_place_name": str}
    PLACE = {"name": str, "description": str}
    THING = {"name": str, "description": str, "start_place_name": str, "end_place_name": str}
    RELATIONSHIP = {"person_1_name": str, "person_2_name": str, "relationship": str}

    All fields are required.

    Important: Only return a single piece of valid JSON text.

    Here is the story:
	%s
`, story)))
	if err != nil {
		log.Printf("extractEntities: %v", err)
		return ""
	}

	ent := string(resp.Candidates[0].Content.Parts[0].(genai.Text))
	ent = strings.ReplaceAll(ent, "```json", "")
	ent = strings.ReplaceAll(ent, "```", "")
	return ent
}
