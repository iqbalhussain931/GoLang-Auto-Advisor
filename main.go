package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type autoAdvisor struct {
	app.Compo
}

func (a *autoAdvisor) onChange(ctx app.Context, e app.Event) {
	files := ctx.JSSrc().Get("files")
	if !files.Truthy() || files.Get("length").Int() == 0 {
		fmt.Println("file not found")
		return
	}

	openFileErrorBox := app.Window().Get("document").Call("getElementById", "openFileErrorBox")

	file := files.Index(0)

	fileType := file.Get("type").String()

	if fileType != "text/plain" {
		openFileErrorBox.Get("classList").Call("remove", "d-none")
		openFileErrorBox.Set("innerHTML", "Invalid file type. Please open file type of (txt).")
		return
	} else {

		fileName := file.Get("name").String()

		var studentName = fileName[:len(fileName)-4]

		openFileErrorBox.Get("classList").Call("add", "d-none")
		openFileErrorBox.Set("innerHTML", "")

		var close func()

		onFileLoad := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			event := args[0]
			content := event.Get("target").Get("result")

			ctx.SetState("fileData", content.String())

			ctx.SetState("studentName", studentName)

			var previewCoursesList []previewCourses

			ctx.SetState("fileData1", previewCoursesList)

			var name []previewCourses
			ctx.GetState("fileData1", &name)

			close()
			return nil
		})

		onFileLoadError := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			// Your error handling...
			openFileErrorBox.Get("classList").Call("remove", "d-none")
			openFileErrorBox.Set("innerHTML", "Unable to open file.")
			close()
			return nil
		})

		// To release resources when callback are called.
		close = func() {
			onFileLoad.Release()
			onFileLoadError.Release()
		}

		reader := app.Window().Get("FileReader").New()
		reader.Set("onload", onFileLoad)
		reader.Set("onerror", onFileLoadError)
		reader.Call("readAsText", file, "UTF-8")
	}

}

func (a *autoAdvisor) openFile(ctx app.Context, e app.Event) {
	e.PreventDefault()
}

func (a *autoAdvisor) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.H1().Body(
				app.Text("Go Language"),
			).Class("w3-xlarge"),
			app.H1().Body(
				app.Text("Auto Advisior"),
			).Class("w3-xlarge"),
			app.Div().Body(
				app.Div().Body(
					app.A().Body(
						app.Text("New Student"),
					).Class("w3-bar-item w3-button").Href("/student-advisor"),
				).Class("w3-bar w3-border"),
			).Class("w3-box"),
			app.Div().Body(
				app.Form().Method("post").OnSubmit(a.openFile).Body(
					app.Div().Body(
						app.Input().Type("file").OnChange(a.onChange).ID("fileInput").Placeholder("Existing Student").Class("w3-bar-item w3-button"),
						app.A().Href("/student-advisor").Body(app.Text("Open File")).ID("fileInput2").Class("w3-bar-item w3-button"),
						app.Div().Body().Class("errorBox d-none col-sm-12").ID("openFileErrorBox"),
					).Class("w3-bar w3-border"),
				).Class("form"),
			).Class("w3-box"),
			app.Div().Body(
				app.Div().Body(
					app.H1().Body(
						app.Text("Group Members"),
					).Class("w3-xlarge"),
					app.Div().Body(
						app.Ul().Body(
							app.Li().Text("Muhammad Iqbal Hussain").Class("member"),
							app.Li().Text("Anthony Mimfa").Class("member"),
							app.Li().Text("Jiahao Deng").Class("member"),
							app.Li().Text("Jing Ma").Class("member"),
							app.Li().Text("Qizheng Ma").Class("member"),
						).Class("list"),
					).Class("group-members-main"),
				).Class("group-members-wrapper"),
			).Class("w3-box"),
		).Class("w3-panel w3-center w3-opacity"),
	).Class("w3-content")
}

func main() {

	app.Route("/", &autoAdvisor{})

	app.Route("/student-advisor", &studentAdvisor{})

	http.HandleFunc("/get-cources", receiveAjax)

	http.HandleFunc("/save-file", downloadFile)

	app.RunWhenOnBrowser()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := app.GenerateStaticWebsite("", &app.Handler{
		Name:               "GoLang Auto Advisor",
		Resources:          app.GitHubPages("GoLang-Auto-Advisor"),
		Author:             "Muhammad Iqbal Hussain",
		BackgroundColor:    "",
		CacheableResources: []string{},
		Description:        "An Auto Advisor developed in Go Language and go-app.",
		Env:                map[string]string{},
		Icon:               app.Icon{},
		Image:              "",
		InternalURLs:       []string{},
		Keywords:           []string{},
		LoadingLabel:       "",
		PreRenderCache:     nil,
		ProxyResources:     []app.ProxyResource{},
		RawHeaders:         []string{},
		Scripts: []string{
			"https://kit.fontawesome.com/8ed9c3141e.js",
			"static/js/bootstrap.min.js",
			"static/js/script.js",
		},
		ShortName: "",
		Styles: []string{
			"static/css/bootstrap.min.css",
			"https://www.w3schools.com/w3css/4/w3.css",
			"https://fonts.googleapis.com/css?family=Raleway",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css",
			"static/css/bootstrap.min.css",
			"static/css/style.css",
		},
		ThemeColor: "",
		Title:      "Go Auto Advisor",
		Version:    "1.0.0",
	})

	http.Handle("/", &app.Handler{})

	if err != nil {
		log.Fatal(err)
	}
}
