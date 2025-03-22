// ui/widgets.go
package ui

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/tags"
)

func MenuWidget() gml.Tag {
	return tags.Div(gml.Text("Menu placeholder"), attr.Class("menu-widget"))
}

func ChatWidget() gml.Tag {
	return tags.Form(gml.Tags(
		tags.Textarea(gml.Empty(), attr.Name("message")),
		tags.Button(gml.Text("Send"), attr.Type("submit")),
	), gml.StringAttribute("method", "post"), gml.StringAttribute("action", "/chat"), attr.Class("chat-widget"))
}
