package controller

import (
	"github.com/Meduzz/yacp/storage"
	"github.com/Meduzz/yacp/ui"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
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
