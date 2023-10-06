package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+"output.txt")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	http.ServeFile(w, r, "/home/student/ascii-art-web/Rupert/output.txt")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		showError(w, "404 Page not found.", 404)
	}

	switch r.Method {
	case "GET":
		//http.ServeFile(w, r, "template/form.html")
		t, err := template.ParseFiles("template/form.html")
		if err == nil {
			t.Execute(w, "")
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Print("HTTP status 500 - Internal Server Errors", err)
			return
		}

		//	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		banner := r.FormValue("banner")
		textbox := r.FormValue("input")
		// fmt.Println("banner is", banner)
		// fmt.Println("textbox is", textbox)
		output, _ := os.Create("output.txt")

		defer output.Close()

		// errorHandler checks for errors, if no error dedected it'll run ascii-art
		if len(banner) == 0 || strings.Contains(textbox, "£") {
			showError(w, "400 Bad Request", 400)
		} else if len(textbox) == 0 {
			showError(w, "500 Internal Server Error", 500)
		} else {

			asciiArt(w, banner, textbox, output)
			ascii, _ := os.ReadFile("output.txt")
			t, err := template.ParseFiles("template/form.html")
			if err == nil {
				t.Execute(w, string(ascii))
			}
			return
		}
	}
}

func showError(w http.ResponseWriter, message string, statusCode int) {
	t, err := template.ParseFiles("template/errors.html")
	if err == nil {
		w.WriteHeader(statusCode)
		t.Execute(w, message)
		return
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/download", downloadHandler)

	fmt.Printf("Starting server at http://localhost:8080 ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP status 500 - Internal Server Errors")
	}
}
