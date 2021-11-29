package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type getCources struct {
	app.Compo
}

var courcesPath = "./static/data/courses.txt"

func (h *getCources) Render() app.UI {

	cources := getCourcesData()

	return app.Select().Body(
		app.Range(cources).Slice(func(i int) app.UI {
			// fmt.Println(cources[i].name.(string))
			// return app.Option().Text("asd")
			return app.Option().Text(cources[i]).Value(cources[i])
		}),
	).Class("form-control").ID("courseDropdown")

	// return app.Option().Text("Iqbal Hussain").Value("Hussain")

}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// ajax_post_data := r.FormValue("ajax_post_data")
		// fmt.Println("Receive ajax post data string ", ajax_post_data)

		// os.Open() opens specific file in
		// read-only mode and this return
		// a pointer of type os.
		file, err := os.Open(courcesPath)

		if err != nil {
			log.Fatalf("failed to open")

		}

		// The bufio.NewScanner() function is called in which the
		// object os.File passed as its parameter and this returns a
		// object bufio.Scanner which is further used on the
		// bufio.Scanner.Split() method.
		scanner := bufio.NewScanner(file)

		// The bufio.ScanLines is used as an
		// input to the method bufio.Scanner.Split()
		// and then the scanning forwards to each
		// new line using the bufio.Scanner.Scan()
		// method.
		scanner.Split(bufio.ScanLines)
		var text []string

		for scanner.Scan() {
			// w.Write([]byte("<h2>" + scanner.Text() + "<h2>"))
			text = append(text, scanner.Text())
		}

		// The method os.File.Close() is called
		// on the os.File object to close the file
		file.Close()

		var responseData []interface{}

		// and then a loop iterates through
		// and prints each of the slice values.
		for _, each_ln := range text {

			singleCourse := make(map[string]interface{})
			// fmt.Println(each_ln)

			course := strings.Split(each_ln, "|")

			if course[0] != "" {
				singleCourse["name"] = course[0]
			} else {
				singleCourse["name"] = ""
			}

			if course[1] != "" {
				singleCourse["credit_hour"] = course[1]
			} else {
				singleCourse["credit_hour"] = ""
			}

			if course[2] != "" {
				singleCourse["pre_req"] = course[2]
			} else {
				singleCourse["pre_req"] = ""
			}

			// fmt.Println(singleCourse)

			responseData = append(responseData, singleCourse)

		}

		jData, err := json.Marshal(responseData)

		if err != nil {
			// handle error
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)

	}
}

type course struct {
	Name        string `json:"name"`
	Credit_hour string `json:"credit_hour"`
	Pre_req     []Prerequisites
}

type Prerequisites struct {
	Cources []course
}

func getYears() []int {

	year, _, _ := time.Now().Date()
	var years []int

	for i := year; i <= 2030; i++ {
		years = append(years, i)
	}

	return years
}

func getSemesters() [3]string {

	var semesters [3]string

	semesters[0] = "Spring"
	semesters[1] = "Summer"
	semesters[2] = "Fall"

	return semesters
}

func getGrades() [11]string {

	var grades [11]string

	grades[0] = "A"
	grades[1] = "A-"
	grades[2] = "B+"
	grades[3] = "B"
	grades[4] = "B-"
	grades[5] = "C+"
	grades[6] = "C"
	grades[7] = "C-"
	grades[8] = "D+"
	grades[9] = "D"
	grades[10] = "F"

	return grades
}

func getCourcesData() []course {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	// file, err := os.Open(courcesPath)

	file, err := http.Get("https://raw.githubusercontent.com/iqbalhussain931/GoLang-Auto-Advisor/main/courses.txt")

	// fmt.Println(resp.Body)

	if err != nil {
		// log.Fatalf("failed to open")
		fmt.Println(err.Error())
	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file.Body)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		// w.Write([]byte("<h2>" + scanner.Text() + "<h2>"))
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Body.Close()

	var responseData []course

	// and then a loop iterates through
	// and prints each of the slice values.
	for _, each_ln := range text {

		singleCourse := course{}

		sigCourse := strings.Split(each_ln, "|")

		if sigCourse[0] != "" {
			singleCourse.Name = sigCourse[0]
		} else {
			singleCourse.Name = ""
		}

		if sigCourse[1] != "" {
			singleCourse.Credit_hour = sigCourse[1]
		} else {
			singleCourse.Credit_hour = ""
		}

		if sigCourse[2] != "" {

			var allPrereqs []Prerequisites

			singlePrereq := Prerequisites{}

			a := regexp.MustCompile(`\s`)
			preReqs := a.Split(sigCourse[2], -1)

			for i, preReq := range preReqs {

				b := regexp.MustCompile(`,`)
				preReqsAndCourses := b.Split(preReq, -1)

				var allSinglePreReqCourse []course

				singlePreReqCourse := course{}

				// Temporary condition need to revisit this logic.
				if preReq == "Senior" {
					singlePreReqCourse.Name = preReqs[i] + " " + preReqs[i+1]
					allSinglePreReqCourse = append(allSinglePreReqCourse, singlePreReqCourse)
					singlePrereq.Cources = allSinglePreReqCourse
					allPrereqs = append(allPrereqs, singlePrereq)
					break
				} else {
					for _, preReqAnd := range preReqsAndCourses {
						singlePreReqCourse.Name = preReqAnd
						allSinglePreReqCourse = append(allSinglePreReqCourse, singlePreReqCourse)
					}
				}

				singlePrereq.Cources = allSinglePreReqCourse

				allPrereqs = append(allPrereqs, singlePrereq)
			}

			singleCourse.Pre_req = allPrereqs
		}

		responseData = append(responseData, singleCourse)

	}

	return responseData
}
