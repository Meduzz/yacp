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
		padding: 2em;
	}

	.menu-widget {
		flex-basis: 25%;
		padding: 2em;
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
func Layout(content ...gml.Tag) gml.Tag {
	return tags.Div(gml.Tags(content...), attr.Class("layout"))
}
