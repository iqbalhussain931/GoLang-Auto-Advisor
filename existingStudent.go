package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type existingStudent struct {
	app.Compo
}

func (h *existingStudent) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.H1().Body(
				app.Text("New Student"),
			).Class("w3-xlarge"),
			app.Template(),
		).Class("w3-panel w3-center w3-opacity"),
	).Class("w3-content")
}
