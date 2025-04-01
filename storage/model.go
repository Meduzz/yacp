package storage

type (
	Chat struct {
		ID       string         // uuid
		Parent   string         // ID of a parent chat
		Name     string         // name
		Host     string         // llm host
		Messages []*ChatMessage // the chat
	}

	ChatMessage struct {
		ID      string // uuid
		Role    string // user|assistant|tool
		Message string // the prompt/response
	}
)
