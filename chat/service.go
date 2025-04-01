package chat

import (
	"encoding/json"
	"fmt"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/yacp/ollama"
	"github.com/Meduzz/yacp/storage"
	"github.com/Meduzz/yacp/tools"
	"github.com/google/uuid"
)

type (
	ChatService struct {
		systemPrompt string
	}

	ToolResult struct {
		Status string          `json:"status"`
		Result json.RawMessage `json:"result"`
	}
)

var (
	service *ChatService
)

func InitChatService(host string) error {
	service = &ChatService{
		systemPrompt: `There's no situation where a Chuck Norris joke will not fit. The responsibility to find or create one is now yours.`,
	}

	return nil
}

func CreateChat(name, host string) (*storage.Chat, error) {
	system := &storage.ChatMessage{
		ID:      uuid.NewString(),
		Role:    "system",
		Message: service.systemPrompt,
	}

	// create a *storage.Chat
	chat := &storage.Chat{
		Name:     name,
		Host:     host,
		ID:       uuid.NewString(),
		Messages: []*storage.ChatMessage{system},
	}

	// call storage.SaveChat
	err := storage.SaveChat(chat)

	if err != nil {
		return nil, err
	}

	return chat, nil
}

func CreateChatMessage(id, message, llm string) (*storage.Chat, error) {
	// load the chat via storage.LoadChat(chat)
	chat, err := storage.LoadChat(id)

	if err != nil {
		return nil, err
	}

	// append a *storage.ChatMessage with the new message to the chat
	chatMessage := &storage.ChatMessage{
		ID:      uuid.NewString(),
		Role:    "user",
		Message: message,
	}

	chat.Messages = append(chat.Messages, chatMessage)

	// send the chat messages of the chat to the llm
	chatRequest := convert(chat, llm)
	chatResponse, err := ollama.Chat(chatRequest)

	if err != nil {
		return nil, err
	}

	// append the response from the llm as a *storage.ChatMessage to the chat
	if chatResponse.Message.Content != "" {
		resp := &storage.ChatMessage{
			ID:      uuid.NewString(),
			Role:    chatResponse.Message.Role,
			Message: chatResponse.Message.Content,
		}

		chat.Messages = append(chat.Messages, resp)
	} else if len(chatResponse.Message.ToolCalls) > 0 {
		toolResults := slice.Map(chatResponse.Message.ToolCalls, func(toolCall *ollama.ToolCall) *ToolResult {
			toolName := toolCall.Function.Name
			println(toolCall.Function.String())

			t := tools.ToolByName(toolName)

			if t == nil {
				return &ToolResult{
					Status: "error",
					Result: json.RawMessage([]byte(fmt.Sprintf("\"tool %s was not found\"", toolName))),
				}
			}

			data, err := t.Execute(toolCall.Function.Arguments)

			if err != nil {
				return &ToolResult{
					Status: "error",
					Result: json.RawMessage([]byte(fmt.Sprintf("\"%s\"", err.Error()))),
				}
			}

			return &ToolResult{
				Status: "success",
				Result: data,
			}
		})

		slice.ForEach(toolResults, func(toolResult *ToolResult) {
			println(toolResult.String())
		})

		bs, err := json.Marshal(toolResults)

		if err != nil {
			return nil, err
		}

		msg := &storage.ChatMessage{
			ID:      uuid.NewString(),
			Role:    "tool",
			Message: string(bs),
		}

		chat.Messages = append(chat.Messages, msg)
	} // TODO else log bad response and return?

	// save the chat via storage.SaveChat(chat)
	err = storage.SaveChat(chat)

	return chat, err
}

func RemoveChatMessage(chatId, msgId string) (*storage.Chat, error) {
	// load the chat via storage.LoadChat(chat)
	chat, err := storage.LoadChat(chatId)

	if err != nil {
		return nil, err
	}

	// remove the *storage.ChatMessage with the matching id from the chat
	chat.Messages = slice.Filter(chat.Messages, func(msg *storage.ChatMessage) bool {
		return msg.ID == msgId
	})

	// save the chat via storage.SaveChat(chat)
	err = storage.SaveChat(chat)

	return chat, err
}

func ListModels() ([]string, error) {
	models, err := ollama.List()

	if err != nil {
		return nil, err
	}

	return slice.Map(models.Models, func(m *ollama.Model) string {
		return m.Model
	}), nil
}

func convert(chat *storage.Chat, llm string) *ollama.ChatRequest {
	req := &ollama.ChatRequest{}

	req.Messages = slice.Map(chat.Messages, func(msg *storage.ChatMessage) *ollama.Message {
		m := &ollama.Message{}
		m.Role = msg.Role
		m.Content = msg.Message
		return m
	})

	slice.ForEach(tools.Tools(), func(t tools.Tool) {
		req.Tools = append(req.Tools, t.Meta())
	})

	req.Model = llm
	req.Stream = false

	return req
}

func (t *ToolResult) String() string {
	return fmt.Sprintf("ToolResult{Status: %s, Result: %s}", t.Status, t.Result)
}
