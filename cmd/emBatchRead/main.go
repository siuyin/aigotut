package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/siuyin/aigotut/emb"
)

func main() {
	f, err := os.Open("embeddings.gob")
	if err != nil {
		log.Fatal(err)
	}

	var orec emb.Rec
	dec := gob.NewDecoder(f)
	for {
		if err := dec.Decode(&orec); err != nil {
			break
		}
		fmt.Println(orec.ID, orec.Title, orec.Embedding[0])
	}
}
