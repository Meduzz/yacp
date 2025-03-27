package ui

import (
	"fmt"
	"strings"

	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/htmx"
	"github.com/Meduzz/gml/logic"
	"github.com/Meduzz/gml/tags"
	"github.com/Meduzz/yacp/storage"
)

func MenuWidget() gml.Tag {
	return tags.Div(gml.Text("Menu placeholder"), attr.Class("menu-widget"))
}

// TODO
// - Render markdown
// - When loading an old chat, there's something with char-encoding (there's becomes thereâ€™s and him, becomes himâ€”)
func ChatWidget(chat *storage.Chat, models []string, prefered string) gml.Tag {
	return tags.Form(gml.Tags(
		tags.Article(logic.Slice(chat.Messages, Bubble(chat.ID), gml.Empty()), attr.Class("messages")),
		tags.Div(gml.Tags(
			tags.Label(gml.Text("Chat:")),
			tags.Textarea(gml.Text(""), attr.Name("message"), gml.AnyAttribute("rows", 5)),
		), attr.Class("input-group")),
		tags.Div(gml.Tags(
			tags.Label(gml.Text("LLM:"), attr.For("llm")),
			tags.Select(gml.Tags(
				logic.Slice(models, ModelOption(prefered), gml.Text("")),
			), attr.Name("llm")),
		), attr.Class("input-group")),
		tags.Div(tags.Button(gml.Text("Send"), attr.Type("submit"))),
	), htmx.Post(fmt.Sprintf("/chat/%s", chat.ID)), attr.Class("chat-widget"), htmx.Target("body"))
}

func ModelOption(prefered string) func(string) gml.Tag {
	return func(name string) gml.Tag {
		if name == prefered {
			return tags.Option(gml.Text(name), attr.Value(name), gml.AnyAttribute("selected", true))
		} else {
			return tags.Option(gml.Text(name), attr.Value(name))
		}
	}
}

func Bubble(chattId string) func(*storage.ChatMessage) gml.Tag {
	return func(message *storage.ChatMessage) gml.Tag {
		if message.Role == "system" {
			// - Hide the system prompt message
			return tags.Details(gml.Tags(gml.Text(message.Message),
				tags.Summary(gml.Tags(gml.Text(message.Role), tags.Button(gml.Text("X"), attr.Type("button"), htmx.Delete(fmt.Sprintf("/message/%s/%s", chattId, message.ID)), htmx.Target("body"), htmx.Trigger("click")))),
			), attr.Class("message"))
		} else if message.Role == "assistant" {
			// - "Hide" thinking
			msg := strings.Replace(message.Message, "<think>", "<details class=\"thinking\">", -1)
			msg = strings.Replace(msg, "</think>", "<summary>Thinking</summary></details>", -1)

			return tags.Div(gml.Tags(
				tags.Span(gml.Text(message.Role)),
				tags.Div(gml.Text(msg)),
				tags.Button(gml.Text("X"), attr.Type("button"), htmx.Delete(fmt.Sprintf("/message/%s/%s", chattId, message.ID)), htmx.Target("body"), htmx.Trigger("click")),
			), attr.Class("message"))
		} else {
			return tags.Div(gml.Tags(
				tags.Span(gml.Text(message.Role)),
				tags.Div(gml.Text(message.Message)),
				tags.Button(gml.Text("X"), attr.Type("button"), htmx.Delete(fmt.Sprintf("/message/%s/%s", chattId, message.ID)), htmx.Target("body"), htmx.Trigger("click")),
			), attr.Class("message"))
		}
	}
}

func ListChats(chats []*storage.Chat) gml.Tag {
	return tags.Div(gml.Tags(
		tags.Ul(logic.Slice(chats, Chat, gml.Text("")), attr.Class("chat-list")),
		tags.A(gml.Text("+"), attr.Href("/new"))))
}

func Chat(chat *storage.Chat) gml.Tag {
	return tags.Li(gml.Tags(tags.A(
		tags.H3(gml.Text(chat.Name)),
		attr.Href(fmt.Sprintf("/chat/%s", chat.ID))),
		tags.Button(gml.Text("X"), attr.Type("button"), htmx.Delete(fmt.Sprintf("/chat/%s", chat.ID)), htmx.Target("body"), htmx.Trigger("click")),
	))
}

func CreateChat() gml.Tag {
	return tags.Form(gml.Tags(
		tags.Div(gml.Tags(
			tags.Label(gml.Text("Name:"), gml.StringAttribute("for", "name")),
			tags.Input(gml.Empty(), attr.Name("name"), attr.Type("text")),
		), attr.Class("input-group")),
		tags.Div(tags.Button(gml.Text("Create"), attr.Type("submit"))),
	), htmx.Post("/new"), htmx.Target("body"))
}
