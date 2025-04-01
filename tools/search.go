package tools

import (
	"encoding/json"
	"fmt"

	"github.com/Meduzz/yacp/ollama"
)

type (
	Search struct{}

	params struct {
		Query string `json:"query"`
	}
)

func (s *Search) Name() string {
	return "search"
}

func (s *Search) Meta() *ollama.Tool {
	return &ollama.Tool{
		Type: ollama.TypeRetrieval,
		Function: &ollama.ToolFunction{
			Name:        "search",
			Description: "Searches for the answer of the universe.",
			Parameters: &ollama.Parameters{
				Type:     ollama.TypeObject,
				Required: []string{"query"},
				Properties: map[string]*ollama.Property{
					"query": {
						Type:        ollama.TypeString,
						Description: "The search query to use.",
					},
				},
			},
		},
	}
}

func (s *Search) Execute(args json.RawMessage) (json.RawMessage, error) {
	p := &params{}
	err := json.Unmarshal(args, p)
	if err != nil {
		return nil, err
	}

	println(p.Query)

	return nil, fmt.Errorf("not implemented yet")
}
