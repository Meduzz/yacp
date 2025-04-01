package ollama

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	ToolType   string
	SchemaType string

	ChatRequest struct {
		Model    string     `json:"model"`
		Stream   bool       `json:"stream"`
		Messages []*Message `json:"messages"`
		Tools    []*Tool    `json:"tools,omitempty"`
	}

	Message struct {
		Role      string      `json:"role"`
		Content   string      `json:"content"`
		ToolCalls []*ToolCall `json:"tool_calls,omitempty"`
	}

	Tool struct {
		Type     ToolType      `json:"type"`
		Function *ToolFunction `json:"function"`
	}

	ToolFunction struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Parameters  *Parameters `json:"parameters,omitempty"`
	}

	Parameters struct {
		Type       SchemaType           `json:"type"`
		Required   []string             `json:"required"`
		Properties map[string]*Property `json:"properties"`
	}

	Property struct {
		Type        SchemaType `json:"type"`
		Description string     `json:"description"`
		Enum        []string   `json:"enum,omitempty"`
	}

	ToolCall struct {
		Function *ToolCallFunction `json:"function"`
	}

	ToolCallFunction struct {
		Index     int             `json:"index,omitempty"`
		Name      string          `json:"name"`
		Arguments json.RawMessage `json:"arguments"`
	}

	ChatResponse struct {
		Model      string    `json:"model"`
		CreatedAt  time.Time `json:"created_at"`
		Message    *Message  `json:"message"`
		DoneReason string    `json:"done_reason,omitempty"`

		Done bool `json:"done"`

		Metrics
	}

	Metrics struct {
		TotalDuration      time.Duration `json:"total_duration,omitempty"`
		LoadDuration       time.Duration `json:"load_duration,omitempty"`
		PromptEvalCount    int           `json:"prompt_eval_count,omitempty"`
		PromptEvalDuration time.Duration `json:"prompt_eval_duration,omitempty"`
		EvalCount          int           `json:"eval_count,omitempty"`
		EvalDuration       time.Duration `json:"eval_duration,omitempty"`
	}

	ListModelResponse struct {
		Models []*Model `json:"models"`
	}

	Model struct {
		Name       string    `json:"name"`
		Model      string    `json:"model"`
		ModifiedAt time.Time `json:"modified_at"`
		Size       int64     `json:"size"`
		Digest     string    `json:"digest"`
	}
)

const (
	TypePrompt    = ToolType("prompt")
	TypeFunction  = ToolType("function")
	TypeRetrieval = ToolType("retrieval")

	TypeString  = SchemaType("string")
	TypeNumber  = SchemaType("number")
	TypeBoolean = SchemaType("boolean")
	TypeArray   = SchemaType("array")
	TypeObject  = SchemaType("object")
	TypeNull    = SchemaType("null")
)

func (t *ToolCallFunction) String() string {
	return fmt.Sprintf("%s(%s)", t.Name, string(t.Arguments))
}
