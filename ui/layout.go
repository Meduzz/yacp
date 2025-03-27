package ui

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/tags"
)

// Page returns the base template for a full page.
func Page(content gml.Tag) gml.Tag {
	style := `
	body {
		align-items: center;
		padding: 2em;
	}

	.layout {
		display: flex;
		flex-direction: column;
		border: 1px solid #d3d3d3;
        border-radius: 8px;
		padding: 2em;
	}

	.chat-widget {
		flex-grow: 1;
		flex-basis: auto;
	}

	.menu-widget {
		flex-basis: 25%;
		padding: 1em;
	}

	.message {
		padding-top: 1em;
		padding-bottom: 1em;
	}

	.message span {
		display: block;
		font-weight: bold;
		margin-bottom: 1em
	}

	.input-group {
		padding-top: 1em;
		padding-bottom: 1em;
	}

	.input-group label {
		display: block;
		font-weight: bold;
	}

	.thinking {
		color: #cccccc;
		padding-top: 1em;
		padding-bottom: 1em;
	}
	`

	return tags.Html(
		gml.Tags(
			tags.Head(gml.Tags(
				tags.Title(gml.Text("YACP")),
				gml.New("style", gml.Text(style)),
				gml.New("script", gml.Text(""), attr.Src("https://unpkg.com/htmx.org@1.9.3")),
			)),
			tags.Body(content),
		),
	)
}

// Layout creates a flex box layout.
func Layout(title string, content ...gml.Tag) gml.Tag {
	return tags.Div(gml.Tags(
		tags.H1(gml.Text(title)),
		gml.Tags(content...),
	), attr.Class("layout"))
}
