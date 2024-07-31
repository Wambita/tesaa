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
	AccountType string
	Email       string
	Password    string
	ConfirmPass string
	Type        string
	Years       string
	License     string
	Kra         string
	Phone       string
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
		AccountType: r.FormValue("account-type"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("pass"),
		ConfirmPass: r.FormValue("confirm"),
		Type:        r.FormValue("type"),
		Years:       r.FormValue("years"),
		License:     r.FormValue("license"),
		Kra:         r.FormValue("kra"),
		Phone:       r.FormValue("phone"),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}
	check_user, _ := http.Get(`http://localhost:3000/users?email=%s` + data.Email)

	// fmt.Println(us)
	if check_user != nil {
		println("Email already exists")
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
	fmt.Println("Institution", instType)
	fmt.Println("Email: ", email)

	// Fetch users data
	url := "http://localhost:3000/users?email=" + email
	res, err := http.Get(url)
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

	// Ensure we have at least one user
	if len(users) == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Find matching user
	isUser := false
	if users[0].Password == password {
		if instType == users[0].InstitutionType {
			isUser = true
			userProfile = users[0]
		}
	}

	if isUser {
		switch instType {
		case "business":
			BusinessHandler(w, r)
		case "microfinance":
			MfiHandler(w, r)
		default:
			http.Error(w, "Unknown institution type", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Invalid credentials or institution type", http.StatusUnauthorized)
	}
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

// business pages routes
func BusinessActiveLoansHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business/active_loans.html"))
	tmpl.Execute(w, nil)
}

func BusinessProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business/business_profile.html"))
	tmpl.Execute(w, nil)
}

func BusinessLoanApplicationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business/loan_applications.html"))
	tmpl.Execute(w, nil)
}

func BusinessTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business/transactions.html"))
	tmpl.Execute(w, nil)
}

func LoanApplicationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business/loan_applications.html"))
	tmpl.Execute(w, nil)
}

func MsiListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/mfi/mfis.html"))
	tmpl.Execute(w, nil)
}

func ApplyLoanHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/business/apply_loan.html"))
	tmpl.Execute(w, nil)
}
