package main

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
)

func main() {
	// For embeddings, use the embedding-001 model
	//client.ModelName = "embedding-001"
	client.ModelName = "text-embedding-004"
	cl := client.New()
	defer cl.Close()

	ctx := context.Background()
	em := cl.Client.EmbeddingModel(client.ModelName)
	res, err := em.EmbedContent(ctx, genai.Text("The quick brown fox jumps over the lazy dog."))
	//res, err := em.EmbedContent(ctx, genai.Text("The sky is blue because of Rayleigh scattering."))

	if err != nil {
		panic(err)
	}
	fmt.Println(res.Embedding.Values)

	s := float32(0.0)
	for _, v := range res.Embedding.Values {
		s += v * v
	}
	fmt.Println("Sum of square of values:", s)

}
