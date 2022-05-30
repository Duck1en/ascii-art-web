package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	ascii "asciiart/ascii"
)

var t *template.Template

type Home struct {
	Input, Output, Font /*err*/ string
}

func main() {
	http.HandleFunc("/", Ascii)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/ascii-art-web", AsciiHandler)

	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func init() {
	t = template.Must(template.ParseGlob("./static/*.html"))
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatalf("template parse error: %v", err)
	}

	tmpl.Execute(w, nil)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-web" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "POST":
		h := Home{}
		h.Output = ""
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		in := r.FormValue("input")
		f := r.FormValue("font")
		f += ".txt"

		h.Input = in
		h.Font = f

		result, err := ascii.AsciiDrawer(in, f)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		h.Output = result
		t.ExecuteTemplate(w, "index.html", h)
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// // fmt.Fprintf(w, "%s", Output)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404: Not Found")
	}
}

// package main

// import (
// 	"artweb/ascii"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"text/template"
// )

// func main() {
// 	HandleRequest()
// }

// var mainTmpl, errTmpl *template.Template

// func HandleRequest() {
// 	var err error
// 	mainTmpl, err = template.ParseFiles("static/index.html")
// 	ErrorCheck(err)
// 	errTmpl, err = template.ParseFiles("static/error.html")
// 	ErrorCheck(err)

// 	fs := http.FileServer(http.Dir("templates"))
// 	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
// 	http.HandleFunc("/", path_page)
// 	http.HandleFunc("/ascii-art", home_page)
// 	fmt.Println("Server is started")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func home_page(page http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		page.WriteHeader(http.StatusMethodNotAllowed)
// 		errTmpl.Execute(page, "405: Status method not allowed")
// 		return
// 	}

// 	font := r.FormValue("banner")
// 	if font != "standard" && font != "shadow" && font != "thinkertoy" {
// 		page.WriteHeader(http.StatusBadRequest)
// 		errTmpl.Execute(page, "400: Bad request")
// 		return
// 	}

// 	text := r.FormValue("data")
// 	if text == "" {
// 		page.WriteHeader(http.StatusBadRequest)
// 		errTmpl.Execute(page, "400: Bad request")
// 		return
// 	}

// 	var textUpd string

// 	for _, v := range text {
// 		if v == 13 {
// 			continue
// 		}
// 		if !(v >= 32 && v <= 126 || v == 10) {
// 			page.WriteHeader(http.StatusBadRequest)
// 			errTmpl.Execute(page, "400: Bad request")
// 			return
// 		}
// 		textUpd += string(v)
// 	}

// 	f, err3 := ascii.AssingFont(font)
// 	ErrorCheck(err3)
// 	finalArt := ascii.OutputAscii(textUpd, f)

// 	btnClicked := r.FormValue("submit_button")

// 	if btnClicked == "Submit" {
// 		mainTmpl.Execute(page, finalArt)
// 	} else {
// 		page.WriteHeader(http.StatusBadRequest)
// 		errTmpl.Execute(page, "400: Bad request")
// 	}
// }

// func path_page(page http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		page.WriteHeader(http.StatusNotFound)
// 		errTmpl.Execute(page, "404: Not found")
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		page.WriteHeader(http.StatusMethodNotAllowed)
// 		errTmpl.Execute(page, "405: Status method not allowed")
// 		return
// 	}
// 	mainTmpl.Execute(page, nil)
// }

// func ErrorCheck(err error) {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(0)
// 	}
// }
