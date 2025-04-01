package tools

import (
	"encoding/json"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/yacp/ollama"
)

type (
	Tool interface {
		Name() string
		Meta() *ollama.Tool
		Execute(args json.RawMessage) (json.RawMessage, error)
	}
)

func Tools() []Tool {
	return []Tool{
		&Search{},
	}
}

func ToolByName(name string) Tool {
	return slice.Head(slice.Filter(Tools(), func(tool Tool) bool {
		return tool.Name() == name
	}))
}
