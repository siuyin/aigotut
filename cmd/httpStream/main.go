package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/generative-ai-go/genai"
	"github.com/siuyin/aigotut/client"
	"github.com/siuyin/aigotut/gfmt"
	"github.com/siuyin/dflt"
)

func main() {
	cl := client.New()
	defer cl.Close()

	genFunc := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		iter := cl.Model.GenerateContentStream(context.Background(),
			//genai.Text("What is the significance of the number 42?"))
			genai.Text("Write a long story, in the sytle of the TV program: 60 Minutes, about the Apollo missions"))
		gfmt.FlusherPrintStreamResponse(w, iter)
	}

	http.HandleFunc("/", genFunc)

	port := dflt.EnvString("PORT", "8080")
	fmt.Println("starting web server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
