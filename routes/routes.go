package routes

import (
	"fmt"
	"net/http"
	"text/template"
)

type RegisterData struct {
	// OrgName     string
	Email       string
	Password    string
	// ConfirmPass string
	// Type        string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/about.html"))
	tmpl.Execute(w, nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("ddd")
	data := RegisterData{
	// 	OrgName:     r.FormValue("orgname"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("pass"),
	// 	ConfirmPass: r.FormValue("confirm"),
		// Type:        r.FormValue("type"),
	}

	 fmt.Print( data.Email, data.Password)

	tmpl := template.Must(template.ParseFiles("template/register.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/login.html"))
	tmpl.Execute(w, nil)
}

func BusinessHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business_dashboard.html"))
	tmpl.Execute(w, nil)
}

func MfiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/mfi_dashboard.html"))
	tmpl.Execute(w, nil)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/error.html"))
	tmpl.Execute(w, nil)
}

func MfiReportHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/mfi/records.html"))
	tmpl.Execute(w, nil)
}

func MfiReportDownloadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/mfi/download.html"))
	tmpl.Execute(w, nil)
}
