package main

import (
	"fmt"
	"net/http"
	"tesaa/routes"
	"text/template"
)

func main() {

	template.ParseGlob("template/*.html")

	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/about", routes.AboutHandler)
	http.HandleFunc("/error", routes.ErrorHandler)
	http.HandleFunc("/register", routes.RegisterHandler)
	http.HandleFunc("/login", routes.LoginHandler)
	http.HandleFunc("/business", routes.BusinessHandler)
	http.HandleFunc("/mfi", routes.MfiHandler)
	http.HandleFunc("/records", routes.MfiReportHandler)
	http.HandleFunc("/download", routes.MfiReportDownloadHandler)

	http.ListenAndServe(":8080", nil)
	fmt.Println("server listening on http://localhost:8080")

}
