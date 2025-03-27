package controller

import (
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/yacp/chat"
	"github.com/Meduzz/yacp/storage"
	"github.com/Meduzz/yacp/ui"
	"github.com/gin-gonic/gin"
)

func RemoveMessage(ctx *gin.Context) {
	chattId := ctx.Param("chat")
	messageId := ctx.Param("message")

	chatt, err := storage.LoadChat(chattId)

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	chatt.Messages = slice.Filter(chatt.Messages, func(msg *storage.ChatMessage) bool {
		return msg.ID != messageId
	})

	err = storage.SaveChat(chatt)

	if err != nil {
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
