// Package client provides functions relating to a Google generative AI client.
package client

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var ModelName = "gemini-1.5-flash"

type Info struct {
	ModelName string
	Client    *genai.Client
	Model     *genai.GenerativeModel
}

func New() *Info {
	inf := &Info{}
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	inf.Client = client
	inf.ModelName = ModelName
	inf.Model = client.GenerativeModel(ModelName)
	return inf
}

func (inf *Info) Close() {
	inf.Client.Close()
}
