// Pacakge gfmt provide formatted output for Google generative AI products.
package gfmt

import (
	"fmt"
	"io"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

func PrintResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}
func FprintResponse(w io.Writer, resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Fprintln(w, part)
			}
		}
	}
}

func PrintStreamResponse(iter *genai.GenerateContentResponseIterator) {
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponse(resp)
	}
}
func FprintStreamResponse(w io.Writer, iter *genai.GenerateContentResponseIterator) {
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		FprintResponse(w, resp)
	}
}
