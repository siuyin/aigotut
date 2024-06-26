// Package emb specified the embedding type format
package emb

type Rec struct {
	ID        string
	Title     string
	Content   string
	Embedding []float32
}
