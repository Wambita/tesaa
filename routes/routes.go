package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

type RegisterData struct {
	// OrgName     string
	Email    string
	Password string
	// ConfirmPass string
	// Type        string
}

type User struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	InstitutionType string `json:"institution_type"`
}

var userProfile User

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
		Email:    r.FormValue("email"),
		Password: r.FormValue("pass"),
		// 	ConfirmPass: r.FormValue("confirm"),
		// Type:        r.FormValue("type"),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}
	resp, err := http.Post("http://localhost:3000/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error posting JSON data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	fmt.Print(data.Email, data.Password)

	tmpl := template.Must(template.ParseFiles("template/register.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/login.html"))
	tmpl.Execute(w, nil)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	instType := r.FormValue("inst-type")

	// Fetch users data
	res, err := http.Get("http://localhost:3000/users")
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	// Read response body
	ResBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal users data
	var users []User
	if err := json.Unmarshal(ResBody, &users); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Find matching user
	isUser := false
	for i := 0; i < len(users); i++ {
		if users[i].Email == email && users[i].Password == password {
			userProfile = users[i]
			if instType == "business" {
				isUser = true
				break
			}
		}
	}

	if isUser {
		BusinessHandler(w, r)
	} else {
		RegisterHandler(w, r)
	}
	// fmt.Println(userProfile)
	// Render template
	// tpl, err := template.ParseFiles("template/auth/auth.html")
	// if err != nil {
	// 	http.Error(w, "Error loading template", http.StatusInternalServerError)
	// 	fmt.Println("Error parsing file:", err)
	// 	return
	// }

	// if err := tpl.Execute(w, userProfile); err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	fmt.Println("Error executing template:", err)
	// }
}

func BusinessHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business_dashboard.html"))
	tmpl.Execute(w, userProfile)
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
