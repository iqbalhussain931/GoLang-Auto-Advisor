package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type studentAdvisor struct {
	app.Compo
	studentName        string
	cources            []course
	pre_req            []Prerequisites
	years              []int
	semesters          [3]string
	grades             [11]string
	selectedCourse     string
	selectedCreditHour string
	selectedPreReqs    []Prerequisites
	selectedYear       int
	selectedSemester   string
	selectedGrade      string
	selectedRow        int
	isFormValid        bool
	previewCourses     []previewCourses
	CGPA               float64
}

// func (h *studentAdvisor) OnMount(ctx app.Context) {
// 	fmt.Println("component Mounted")

// 	h.cources = getCourcesData()

// 	elem := app.Window().Get("document").Call("getElementById", "courseDropdown")

// 	elem.Call("append", app.Raw("<option>Iqbal Huusain</option>"))

// 	fmt.Println("Hello")
// 	h.Update()
// }

type previewCourses struct {
	Course          string `json:"courseName"`
	Credit_hour     string `json:"creditHour"`
	Grade           string `json:"grade"`
	Prerequisites   []Prerequisites
	Semester_n_year string `json:"semesterNYear"`
	Semester        string `json:"semester"`
	Year            int    `json:"year"`
}

func (h *studentAdvisor) validateFields(app app.BrowserWindow) {

	studentNameErrorBox := app.Get("document").Call("getElementById", "studentNameErrorBox")
	courseErrorBox := app.Get("document").Call("getElementById", "courseErrorBox")
	yearErrorBox := app.Get("document").Call("getElementById", "yearErrorBox")
	semesterErrorBox := app.Get("document").Call("getElementById", "semesterErrorBox")
	gradeErrorBox := app.Get("document").Call("getElementById", "gradeErrorBox")

	isValid := false

	if h.studentName == "" {
		studentNameErrorBox.Get("classList").Call("remove", "d-none")
		studentNameErrorBox.Set("innerHTML", "Please enter student name")
		isValid = false
	} else {
		studentNameErrorBox.Get("classList").Call("add", "d-none")
		studentNameErrorBox.Set("innerHTML", "")
		isValid = true
	}

	if h.selectedCourse == "" || h.selectedCourse == "none" {
		courseErrorBox.Get("classList").Call("remove", "d-none")
		courseErrorBox.Set("innerHTML", "Please select a course")
		isValid = false
	} else {
		courseErrorBox.Get("classList").Call("add", "d-none")
		courseErrorBox.Set("innerHTML", "")
		isValid = true
	}

	if h.selectedYear == 0 {
		yearErrorBox.Get("classList").Call("remove", "d-none")
		yearErrorBox.Set("innerHTML", "Please select a year")
		isValid = false
	} else {
		yearErrorBox.Get("classList").Call("add", "d-none")
		yearErrorBox.Set("innerHTML", "")
		isValid = true
	}

	if h.selectedSemester == "" || h.selectedSemester == "none" {
		semesterErrorBox.Get("classList").Call("remove", "d-none")
		semesterErrorBox.Set("innerHTML", "Please select a semester")
		isValid = false
	} else {
		semesterErrorBox.Get("classList").Call("add", "d-none")
		semesterErrorBox.Set("innerHTML", "")
		isValid = true
	}

	if h.selectedGrade == "" || h.selectedGrade == "none" {
		gradeErrorBox.Get("classList").Call("remove", "d-none")
		gradeErrorBox.Set("innerHTML", "Please select a grade")
		isValid = false
	} else {
		gradeErrorBox.Get("classList").Call("add", "d-none")
		gradeErrorBox.Set("innerHTML", "")
		isValid = true
	}

	// if len(h.previewCourses) > 0 {

	// 	preReqNotSatisfied := true
	// 	preReqNotSatisfiedCourse := ""

	// 	for _, course := range h.previewCourses {

	// 		// fmt.Println(1, course.Course)

	// 		for _, req := range h.selectedPreReqs {

	// 			// fmt.Println(2, req.Cources)

	// 			if req.Cources == course.Course {
	// 				preReqNotSatisfied = false
	// 				break
	// 			} else {
	// 				preReqNotSatisfiedCourse = req.Cources
	// 			}
	// 		}

	// 		if !preReqNotSatisfied {
	// 			break
	// 		}

	// 	}

	// 	if preReqNotSatisfied {
	// 		courseErrorBox.Get("classList").Call("remove", "d-none")
	// 		courseErrorBox.Set("innerHTML", "The prerequisites requirement is not meet. Please add "+preReqNotSatisfiedCourse+" first.")
	// 		isValid = false
	// 	} else {
	// 		isValid = true
	// 	}

	// }

	h.isFormValid = isValid
}

func (h *studentAdvisor) onCourseChange(ctx app.Context, e app.Event) {

	h.selectedCourse = ctx.JSSrc().Get("value").String()

	if h.selectedCourse == "none" {

		h.selectedCreditHour = "0"
		h.pre_req = []Prerequisites{}
		h.isFormValid = false

	} else {
		for _, element := range h.cources {

			if element.name == h.selectedCourse {
				h.selectedCreditHour = element.credit_hour
				h.selectedPreReqs = element.pre_req
			}

		}
	}

	h.validateFields(app.Window())

	h.Update()
}

func (h *studentAdvisor) onStudentNameInputChange(ctx app.Context, e app.Event) {
	h.studentName = ctx.JSSrc().Get("value").String()

	h.validateFields(app.Window())
}

func (h *studentAdvisor) onYearInputChange(ctx app.Context, e app.Event) {
	v := ctx.JSSrc().Get("value").String()
	h.selectedYear, _ = strconv.Atoi(v)
	h.validateFields(app.Window())
}

func (h *studentAdvisor) onSemesterInputChange(ctx app.Context, e app.Event) {
	h.selectedSemester = ctx.JSSrc().Get("value").String()
	h.validateFields(app.Window())
}

func (h *studentAdvisor) onGradeInputChange(ctx app.Context, e app.Event) {
	h.selectedGrade = ctx.JSSrc().Get("value").String()
	h.validateFields(app.Window())
}

func (h *studentAdvisor) onEditButtonClick(ctx app.Context, e app.Event) {

	v := ctx.JSSrc().Call("getAttribute", "data-row-id").String()
	h.selectedRow, _ = strconv.Atoi(v)

	for key, element := range h.previewCourses {

		if key == h.selectedRow {
			h.selectedCourse = element.Course
			h.selectedCreditHour = element.Credit_hour
			h.selectedPreReqs = element.Prerequisites
			h.selectedGrade = element.Grade
			h.selectedSemester = element.Semester
			h.selectedYear = element.Year
		}
	}

	h.calculateCGPA(ctx)

	h.Update()

}

func (h *studentAdvisor) onDeleteButtonClick(ctx app.Context, e app.Event) {

	v := ctx.JSSrc().Call("getAttribute", "data-row-id").String()
	h.selectedRow, _ = strconv.Atoi(v)

	for key, _ := range h.previewCourses {

		if key == h.selectedRow {

			h.previewCourses = append(h.previewCourses[:key], h.previewCourses[key+1:]...)

		}

	}

	h.calculateCGPA(ctx)

	h.Update()

}

func (h *studentAdvisor) calculateCGPA(ctx app.Context) {

	type GradePoints struct {
		grade string
		point float64
	}

	var gradePoints []GradePoints
	var points float64

	if len(h.previewCourses) > 0 {

		gradePoints = append(gradePoints, GradePoints{grade: "A", point: 4.00})
		gradePoints = append(gradePoints, GradePoints{grade: "A-", point: 3.67})
		gradePoints = append(gradePoints, GradePoints{grade: "B+", point: 3.33})
		gradePoints = append(gradePoints, GradePoints{grade: "B", point: 3.00})
		gradePoints = append(gradePoints, GradePoints{grade: "B-", point: 2.67})
		gradePoints = append(gradePoints, GradePoints{grade: "C+", point: 2.33})
		gradePoints = append(gradePoints, GradePoints{grade: "C", point: 2.00})
		gradePoints = append(gradePoints, GradePoints{grade: "C-", point: 1.67})
		gradePoints = append(gradePoints, GradePoints{grade: "D+", point: 1.33})
		gradePoints = append(gradePoints, GradePoints{grade: "D", point: 1.00})
		gradePoints = append(gradePoints, GradePoints{grade: "F", point: 0.00})

		for _, course := range h.previewCourses {

			for _, grade := range gradePoints {

				if course.Grade == grade.grade {
					points += grade.point
				}

			}
		}

		h.CGPA = points / float64(len(h.previewCourses))

	} else {
		h.CGPA = 0.00
	}
}

func (h *studentAdvisor) OnSubmitForm(ctx app.Context, e app.Event) {
	e.PreventDefault()

	h.validateFields(app.Window())

	if h.isFormValid {

		notInPreview := true

		for key, course := range h.previewCourses {

			if course.Course == h.selectedCourse {
				h.previewCourses[key].Course = h.selectedCourse
				h.previewCourses[key].Credit_hour = h.selectedCreditHour
				h.previewCourses[key].Prerequisites = h.selectedPreReqs
				h.previewCourses[key].Semester_n_year = h.selectedSemester + " " + strconv.Itoa(h.selectedYear)
				h.previewCourses[key].Grade = h.selectedGrade
				h.previewCourses[key].Semester = h.selectedSemester
				h.previewCourses[key].Year = h.selectedYear
				notInPreview = false
			}

		}

		if notInPreview {
			course := previewCourses{}

			course.Course = h.selectedCourse
			course.Credit_hour = h.selectedCreditHour
			course.Prerequisites = h.selectedPreReqs
			course.Semester_n_year = h.selectedSemester + " " + strconv.Itoa(h.selectedYear)
			course.Grade = h.selectedGrade
			course.Semester = h.selectedSemester
			course.Year = h.selectedYear

			h.previewCourses = append(h.previewCourses, course)
		}

		// Remove Selected class from preview row
		// var selectedRow = app.Window().Get("document").Call("querySelector", ".single-preview.selected")
		// fmt.Println(app.ValueOf(selectedRow))
		// if app.ValueOf(selectedRow) != app.Undefined() {
		// 	selectedRow.Get("classList").Call("remove", "selected")
		// }

		// temp := js.Global().Get()

		// for i, j := range temp {

		// }

		// fmt.Println(temp[0])

		h.calculateCGPA(ctx)

		h.Update()
	}
}

func (h *studentAdvisor) OnMount(ctx app.Context) {
	h.cources = getCourcesData()
	h.years = getYears()
	h.semesters = getSemesters()
	h.grades = getGrades()
	h.isFormValid = false
	var name string
	ctx.GetState("fileData", &name)

	fmt.Println(2)
	fmt.Println(name)
}

func (h *studentAdvisor) Render() app.UI {

	return app.Div().Body(
		app.Header().Body(
			app.Div().Body(
				app.Div().Body(
					app.Div().Body(
						app.Div().Body(
							app.H5().Body(
								app.Text("SOUTHEAST MISSOURI STATE UNIVERSITY"),
							).Class("card-title h5"),
							app.H5().Body(
								app.Text("CS-609 - GoLang Auto Advisor"),
							).Class("card-title h5"),
							app.P().Body(
								app.Text("Dr. Robert Lowe"),
							).Class("card-text"),
						).Class("card-body"),
					).Class("card"),
				).Class("col"),
			).Class("row"),
		).Class("container-fluid w3-center w3-opacity"),
		app.Div().Body(
			app.Div().Body(
				app.Div().Body(
					app.Form().Method("post").OnSubmit(h.OnSubmitForm).Body(
						app.Div().Body(
							app.Label().Body(
								app.Text("Student Name"),
							).Class("form-label col-form-label col-sm-4"),
							app.Div().Body(
								app.Input().Type("text").OnChange(h.onStudentNameInputChange).Placeholder("Enter Student Name").Class("form-control"),
							).Class("col-sm-8"),
							app.Div().Body().Class("errorBox d-none col-sm-8").ID("studentNameErrorBox"),
						).Class("mb-3 row"),
						app.Div().Body(
							app.Label().Body(
								app.Text("Course"),
							).Class("form-label col-form-label col-sm-4"),
							app.Div().Body(
								app.Select().OnChange(h.onCourseChange).Body(
									app.Option().Text("Select Course").Selected(true).Value("none"),
									app.Range(h.cources).Slice(func(i int) app.UI {
										return app.Option().Text(h.cources[i].name).Value(h.cources[i].name)
									}),
								).Class("form-control").ID("courseDropdown"),
							).Class("col-sm-8"),
							app.Div().Body().Class("errorBox d-none col-sm-8").ID("courseErrorBox"),
						).Class("mb-3 row"),
						app.Div().Body(
							app.Label().Body(
								app.Text("Credit Hour"),
							).Class("form-label col-form-label col-sm-4"),
							app.Div().Body(
								app.Input().Type("text").Disabled(true).Placeholder("0").Value(h.selectedCreditHour).Class("form-control"),
							).Class("col-sm-8"),
						).Class("mb-3 row"),
						app.Div().Body(
							app.Label().Body(
								app.Text("Course Prerequisites"),
							).Class("form-label col-form-label col-sm-5"),
							app.Div().Body(
								app.If(len(h.selectedPreReqs) == 0,
									app.Input().Type("text").Disabled(true).Placeholder("None").Value("None").Class("form-control"),
								).Else(
									app.Range(h.selectedPreReqs).Slice(func(i int) app.UI {
										if i != 0 {
											return app.Div().Body(
												app.Span().Text("OR"),
												app.Input().Type("text").Disabled(true).Placeholder("None").Value(h.selectedPreReqs[i].Cources).Class("form-control"),
											).Class("singlePreReq")
										} else {
											return app.Input().Type("text").Disabled(true).Placeholder("None").Value(h.selectedPreReqs[i].Cources).Class("form-control")
										}
									}),
								),
							).Class("col-sm-7"),
						).Class("mb-3 row"),
						app.Div().Body(
							app.Div().Body(
								app.Div().Body(
									app.Label().Body(
										app.Text("Year"),
									).Class("form-label col-form-label col-sm-4"),
									app.Div().Body(
										app.Select().OnChange(h.onYearInputChange).Body(
											app.Option().Text("Select Year").Selected(true).Value(0),
											app.Range(h.years).Slice(func(i int) app.UI {
												if h.years[i] == h.selectedYear {
													return app.Option().Selected(true).Text(h.years[i]).Value(strconv.Itoa(h.years[i]))
												} else {
													return app.Option().Text(h.years[i]).Value(strconv.Itoa(h.years[i]))
												}
											}),
										).Class("form-control").ID("yearDropdown"),
									).Class("col-sm-8"),
									app.Div().Body().Class("errorBox d-none col-sm-8").ID("yearErrorBox"),
								).Class("row"),
							).Class("col-5"),
							app.Div().Body(
								app.Div().Body(
									app.Label().Body(
										app.Text("Semester"),
									).Class("form-label col-form-label col-sm-4"),
									app.Div().Body(
										app.Select().OnChange(h.onSemesterInputChange).Body(
											app.Option().Text("Select Semester").Selected(true).Value("none"),
											app.Range(h.semesters).Slice(func(i int) app.UI {
												if h.semesters[i] == h.selectedSemester && h.selectedSemester != "" {
													return app.Option().Selected(true).Text(h.semesters[i]).Value(h.semesters[i])
												} else {
													return app.Option().Text(h.semesters[i]).Value(h.semesters[i])
												}
											}),
										).Class("form-control").ID("yearDropdown"),
									).Class("col-sm-8"),
									app.Div().Body().Class("errorBox d-none col-sm-8").ID("semesterErrorBox"),
								).Class("row"),
							).Class("col-7"),
							app.Div().Body(
								app.Div().Body(
									app.Div().Body(
										app.Label().Body(
											app.Text("Expected Grade"),
										).Class("form-label col-form-label col-sm-6"),
										app.Div().Body(
											app.Select().OnChange(h.onGradeInputChange).Body(
												app.Option().Text("Select Grade").Selected(true).Value("none"),
												app.Range(h.grades).Slice(func(i int) app.UI {
													if h.grades[i] == h.selectedGrade && h.selectedGrade != "" {
														return app.Option().Selected(true).Text(h.grades[i]).Value(h.grades[i])
													} else {
														return app.Option().Text(h.grades[i]).Value(h.grades[i])
													}
												}),
											).Class("form-control").ID("yearDropdown"),
										).Class("col-sm-6"),
										app.Div().Body().Class("errorBox d-none col-sm-8").ID("gradeErrorBox"),
									).Class("row").Style("margin-top", "15px"),
								).Class("row"),
							).Class("col-8"),
						).Class("mb-3 row"),
						app.Div().Body(
							app.Div().Body(
								// app.A().Target("_blank").Href("/save-file").Body(app.Text("Save File")).Class("w3-bar-item w3-button btn btn-success"),
								app.Button().OnClick(h.onSaveFile).Type("button").Body(app.Text("Save File")).Class("w3-bar-item w3-button btn btn-success"),
							).Class("col"),
							app.Div().Body(
								app.Div().Body().Class("errorBox d-none col-sm-8").ID("saveFileErrorBox"),
							).Class("col"),
							app.Div().Body(
								app.Button().Type("submit").Body(app.Text("Add Course")).Class("w3-bar-item w3-button btn btn-primary"),
							).Class("col"),
						).Class("mb-3 row form-buttons"),
					).Class("form"),
				).Class("col-5"),
				app.Div().Body(
					app.Div().Body(
						app.Div().Body(
							app.H3().Body(
								app.Text("Preview Degree"),
							).Class("card-title heading"),
						).Class("col"),
						app.Div().Body().Class("col"),
						app.Label().Body(
							app.Text("CGPA"),
						).Class("form-label col-form-label col-sm-2"),
						app.Div().Body(
							app.If(
								math.Floor(h.CGPA*100)/100 > 2.50,
								app.Input().Type("text").Disabled(true).Value(fmt.Sprintf("%.2f", math.Floor(h.CGPA*100)/100)).Class("form-control good-gpa"),
							).ElseIf(
								len(h.previewCourses) == 0 && math.Floor(h.CGPA*100)/100 == 0.00,
								app.Input().Type("text").Disabled(true).Value(fmt.Sprintf("%.2f", math.Floor(h.CGPA*100)/100)).Class("form-control"),
							).Else(
								app.Input().Type("text").Disabled(true).Value(fmt.Sprintf("%.2f", math.Floor(h.CGPA*100)/100)).Class("form-control bad-gpa"),
							),
						).Class("col-sm-2"),
					).Class("mb-3 row"),
					app.Div().Body(
						app.Div().Body(
							app.Table().Body(
								app.THead().Body(
									app.Tr().Body(
										app.Th().Body().Text("Course"),
										app.Th().Body().Text("Credit Hours"),
										app.Th().Body().Text("Prerequisites"),
										app.Th().Body().Text("Grade"),
										app.Th().Body().Text("Planned Semester"),
										app.Th().Body().Text("Actions"),
									),
								),
								app.TBody().Body(
									app.If(
										len(h.previewCourses) == 0,
										app.Tr().Body(
											app.Text("No Course Added"),
										).Class("no-preview"),
									).Else(
										app.Range(h.previewCourses).Slice(func(i int) app.UI {
											return app.Tr().Body(
												app.Td().Body().Text(h.previewCourses[i].Course).Class("align-middle"),
												app.Td().Body().Text(h.previewCourses[i].Credit_hour).Class("align-middle"),
												app.If(len(h.previewCourses[i].Prerequisites) == 0,
													app.Td().Body().Text("None").Class("align-middle"),
												).Else(
													app.Td().Body(
														app.Range(h.previewCourses[i].Prerequisites).Slice(func(j int) app.UI {
															if j != 0 {
																return app.Div().Body(
																	app.Strong().Body(
																		app.Text("OR"),
																	),
																	app.Span().Body(
																		app.Text(h.previewCourses[i].Prerequisites[j].Cources),
																	).Class("small-text"),
																)
															} else {
																return app.Span().Body(
																	app.Text(h.previewCourses[i].Prerequisites[j].Cources),
																).Class("normal-text")
															}
														}),
													).Class("align-middle"),
												),
												app.Td().Body().Text(h.previewCourses[i].Grade).Class("align-middle"),
												app.Td().Body().Text(h.previewCourses[i].Semester_n_year).Class("align-middle"),
												app.Td().Body(
													app.Button().OnClick(h.onEditButtonClick).Body(
														app.I().Class("far fa-edit"),
													).Attr("data-row-id", i).Attr("data-bs-toggle", "tooltip").Attr("data-bs-placement", "top").Attr("title", "Edit Row").Class("btn").ID("btnEdit"),
													app.Button().OnClick(h.onDeleteButtonClick).Body(
														app.I().Class("fas fa-trash"),
													).Attr("data-row-id", i).Attr("data-bs-toggle", "tooltip").Attr("data-bs-placement", "top").Attr("title", "Delete Row").Class("btn").ID("btnDelete"),
												).Class("align-middle"),
											).Class("single-preview justify-content-between")
										}),
									),
								),
							).Class("table"),
						).Class("previews"),
					).Class("mb-3"),
				).Class("col-7").ID("previewWrapper"),
			).Class("row"),
		).Class("main-wrapper container-fluid w3-opacity"),
	).Class("container")
}
