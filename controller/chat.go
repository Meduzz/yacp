package controller

import (
	"errors"

	"github.com/Meduzz/yacp/chat"
	"github.com/Meduzz/yacp/storage"
	"github.com/Meduzz/yacp/ui"
	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
)

func ShowChat(ctx *gin.Context) {
	chatId := ctx.Param("chat")
	chatt, err := storage.LoadChat(chatId)

	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			ctx.Redirect(303, "/")
			return
		}

		ctx.AbortWithError(500, err)
		return
	}

	models, err := chat.ListModels()

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	chatWidget := ui.ChatWidget(chatt, models, "")
	layout := ui.Layout(chatt.Name, chatWidget)
	page := ui.Page(layout)

	ui.Render(ctx, page)
}

func HandleChat(ctx *gin.Context) {
	var chatt *storage.Chat
	var err error
	chatId := ctx.Param("chat")

	message, messageOk := ctx.GetPostForm("message")
	llm, llmOk := ctx.GetPostForm("llm")

	if messageOk && llmOk {
		chatt, err = chat.CreateChatMessage(chatId, message, llm)
	} else {
		chatt, err = storage.LoadChat(chatId)
	}

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	models, err := chat.ListModels()

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	chatWidget := ui.ChatWidget(chatt, models, llm)
	layout := ui.Layout(chatt.Name, chatWidget)
	page := ui.Page(layout)

	ui.Render(ctx, page)
}

func RemoveChat(ctx *gin.Context) {
	chattId := ctx.Param("chat")

	err := storage.RemoveChat(chattId)

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	chats, err := storage.ListChats()

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	chatsWidget := ui.ListChats(chats)
	layoutWidget := ui.Layout("Previous chats", chatsWidget)
	page := ui.Page(layoutWidget)

	ui.Render(ctx, page)
}
