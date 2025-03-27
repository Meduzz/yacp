package chat

import (
	"context"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/yacp/storage"
	"github.com/google/uuid"
	"github.com/ollama/ollama/api"
)

type ChatService struct {
	client       *api.Client
	systemPrompt string
}

var (
	service *ChatService
)

func InitChatService(host string) error {
	client, err := api.ClientFromEnvironment()

	if err != nil {
		return err
	}

	service = &ChatService{
		client:       client,
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
		Prompt:   service.systemPrompt,
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
	ctx := context.Background()
	streaming := false
	chatRequest := &api.ChatRequest{
		Model:  llm,
		Stream: &streaming,
		Messages: slice.Map(chat.Messages, func(msg *storage.ChatMessage) api.Message {
			return api.Message{
				Role:    msg.Role,
				Content: msg.Message,
			}
		}),
	}

	responseChannel := make(chan api.ChatResponse, 10)

	err = service.client.Chat(ctx, chatRequest, func(resp api.ChatResponse) error {
		responseChannel <- resp
		return nil
	})

	if err != nil {
		return nil, err
	}

	response := <-responseChannel

	// append the response from the llm as a *storage.ChatMessage to the chat
	resp := &storage.ChatMessage{
		ID:      uuid.NewString(),
		Role:    response.Message.Role,
		Message: response.Message.Content,
	}

	chat.Messages = append(chat.Messages, resp)

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
	models, err := service.client.List(context.Background())

	if err != nil {
		return nil, err
	}

	return slice.Map(models.Models, func(m api.ListModelResponse) string {
		return m.Model
	}), nil
}
