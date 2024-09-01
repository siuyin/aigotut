// Pacakge gfmt provide formatted output for Google generative AI products.
package gfmt

import (
	"fmt"
	"io"
	"log"
	"net/http"

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
func FlusherPrintStreamResponse(w http.ResponseWriter, iter *genai.GenerateContentResponseIterator) {
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		FlusherPrintResponse(w, resp)
	}
}

func FlusherPrintResponse(w http.ResponseWriter, resp *genai.GenerateContentResponse) {
	f, _ := w.(http.Flusher)
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				w.Write([]byte(part.(genai.Text)))
				f.Flush()
			}
		}
	}
}
