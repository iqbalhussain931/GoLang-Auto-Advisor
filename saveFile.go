package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		var courses []previewCourses
		json.Unmarshal([]byte(r.Form.Get("courses")), &courses)

		var fileName = "./" + r.Form.Get("studentName") + ".txt"
		var content string

		for _, course := range courses {
			content += course.Course + "|" + course.Credit_hour + "|"
			if len(course.Prerequisites) > 0 {
				for i := range course.Prerequisites {
					if i == len(course.Prerequisites)-1 {
						for j, preq := range course.Prerequisites[i].Cources {
							if j == len(course.Prerequisites[i].Cources)-1 {
								content += preq.Name + "|"
							} else {
								content += preq.Name + ","
							}
						}

					} else {
						for j, preq := range course.Prerequisites[i].Cources {
							if j == len(course.Prerequisites[i].Cources)-1 {
								content += preq.Name + " "
							} else {
								content += preq.Name + ","
							}
						}
					}
				}
			} else {
				content += "|"
			}

			content += course.Semester_n_year + "|" + course.Grade + "\n"
		}

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		f.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write([]byte(content))
		if err != nil {
			log.Fatal(err)
		}

		f.Close()

		// w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		// w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	}
}

func (h *studentAdvisor) onSaveFile(ctx app.Context, e app.Event) {

	saveFileErrorBox := app.Window().Get("document").Call("getElementById", "saveFileErrorBox")

	if len(h.previewCourses) > 0 {
		cources, _ := json.Marshal(h.previewCourses)

		formData := url.Values{
			"studentName": {h.studentName},
			"courses":     {string(cources)},
		}
		_, err := http.PostForm("/save-file", formData)

		if err != nil {
			saveFileErrorBox.Get("classList").Call("remove", "d-none")
			saveFileErrorBox.Get("classList").Call("remove", "success")
			saveFileErrorBox.Set("innerHTML", "File is not save.")
		} else {
			saveFileErrorBox.Get("classList").Call("remove", "d-none")
			saveFileErrorBox.Get("classList").Call("add", "success")
			saveFileErrorBox.Set("innerHTML", "File is save.")
		}
	} else {
		saveFileErrorBox.Get("classList").Call("remove", "d-none")
		saveFileErrorBox.Get("classList").Call("remove", "success")
		saveFileErrorBox.Set("innerHTML", "Please add courses first.")
	}

	time.AfterFunc(2*time.Second, func() {
		saveFileErrorBox.Get("classList").Call("add", "d-none")
		saveFileErrorBox.Set("innerHTML", "")
	})
}
