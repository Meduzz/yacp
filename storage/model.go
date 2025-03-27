package storage

type (
	Chat struct {
		ID       string         // uuid
		Parent   string         // ID of a parent chat
		Name     string         // name
		Prompt   string         // system prompt
		Host     string         // llm host
		Messages []*ChatMessage // the chat
	}

	ChatMessage struct {
		ID      string // uuid
		Role    string // user|assistant
		Message string // the prompt/response
	}
)
