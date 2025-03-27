package controller

import (
	"fmt"

	"github.com/Meduzz/yacp/chat"
	"github.com/Meduzz/yacp/ui"
	"github.com/gin-gonic/gin"
)

func CreateForm(ctx *gin.Context) {
	form := ui.CreateChat()
	layout := ui.Layout("Create prompt", form)
	page := ui.Page(layout)

	ui.Render(ctx, page)
}

func CreateHandler(ctx *gin.Context) {
	name, nameOk := ctx.GetPostForm("name")

	if !nameOk {
		ctx.AbortWithError(400, fmt.Errorf("name are required values"))
		return
	}

	chatt, err := chat.CreateChat(name, "")

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	models, err := chat.ListModels()

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.Header("HX-Replace-Url", fmt.Sprintf("/chat/%s", chatt.ID))

	chatWidget := ui.ChatWidget(chatt, models, "")
	layout := ui.Layout(chatt.Name, chatWidget)
	page := ui.Page(layout)

	ui.Render(ctx, page)
}
