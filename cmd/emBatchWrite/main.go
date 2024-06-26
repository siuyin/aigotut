package main

import (
	"context"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
	"github.com/siuyin/aigotut/emb"
)

func main() {
	// For embeddings, use the embedding-001 model
	//client.ModelName = "embedding-001"
	client.ModelName = "text-embedding-004"
	cl := client.New()
	defer cl.Close()

	// fields: 0:ID, 1:Title, 2:Content
	f, err := os.Open("./input.csv")
	r := csv.NewReader(f)
	recs, err := r.ReadAll()

	ctx := context.Background()
	em := cl.Client.EmbeddingModel(client.ModelName)
	b := em.NewBatch()
	for i, v := range recs {
		if i == 0 {
			continue
		}
		b.AddContentWithTitle(v[1], genai.Text(v[2]))
	}

	res, err := em.BatchEmbedContents(ctx, b)
	if err != nil {
		log.Fatal(err)
	}

	ofile, err := os.Create("embeddings.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer ofile.Close()

	en := gob.NewEncoder(ofile)
	for i := range recs {
		if i == 0 {
			continue
		}
		if err := en.Encode(emb.Rec{ID: recs[i][0], Title: recs[i][1], Content: recs[i][2], Embedding: res.Embeddings[i-1].Values}); err != nil {
			log.Fatal(err)
		}
		fmt.Println("gob:", i)
	}
	//for _, e := range res.Embeddings {
	//	fmt.Println(e.Values)
	//	fmt.Println("++++++++++++++")
	//}

	fmt.Println("embeddings written to gob file")
}
