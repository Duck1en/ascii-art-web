package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	asciiart "asciiart/ascii"
)

var mainTmpl, errTmpl *template.Template

var (
	ErrNotFound    = errors.New("page does not exist")
	ErrBadRequest  = errors.New("bad request")
	ErrEmptyInput  = errors.New("no input given")
	ErrWrongMethod = errors.New("method not allowed")
)

var port = ":8080"

func main() {
	var err error
	mainTmpl, err = template.ParseFiles("static/index.html")
	ErrorCheck(err)
	errTmpl, err = template.ParseFiles("static/error.html")
	ErrorCheck(err)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/", path_page)
	http.HandleFunc("/ascii-art-web", home_page)
	fmt.Printf("Starting server at port %s\nhttp://localhost%s/\n", port, port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func home_page(page http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(page, http.StatusMethodNotAllowed, ErrWrongMethod)
		return
	}

	font := r.FormValue("banner")
	text := r.FormValue("input")
	color := r.FormValue("color")
	if text == "" {
		errorHandler(page, http.StatusBadRequest, ErrBadRequest)
		return
	}

	finalArt, code, err := asciiart.AsciiDrawer(text, font)
	if err != nil {
		errorHandler(page, code, err)
		return
	}

	d := struct {
		Data  string
		Color string
	}{
		Data:  finalArt,
		Color: color,
	}

	mainTmpl.Execute(page, d)
}

func path_page(page http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(page, http.StatusNotFound, ErrNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(page, http.StatusBadRequest, ErrBadRequest)
		return
	}
	mainTmpl.Execute(page, nil)
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func errorHandler(page http.ResponseWriter, status int, err error) {
	page.WriteHeader(status)
	errorText := fmt.Sprintf("%d %s\n%v", status, http.StatusText(status), err)
	errTmpl.Execute(page, errorText)
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"text/template"

// 	ascii "asciiart/ascii"
// )

// var t *template.Template

// type Home struct {
// 	Input, Output, Font /*err*/ string
// }

// func main() {
// 	http.HandleFunc("/", Ascii)
// 	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
// 	http.HandleFunc("/ascii-art-web", AsciiHandler)

// 	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func init() {
// 	t = template.Must(template.ParseGlob("./static/*.html"))
// }

// func Ascii(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		errorHandler(w, r, http.StatusNotFound)
// 		return
// 	}
// 	tmpl, err := template.ParseFiles("static/index.html")
// 	if err != nil {
// 		log.Fatalf("template parse error: %v", err)
// 	}

// 	tmpl.Execute(w, nil)
// }

// func AsciiHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/ascii-art-web" {
// 		errorHandler(w, r, http.StatusNotFound)
// 		return
// 	}
// 	switch r.Method {
// 	case "POST":
// 		h := Home{}
// 		h.Output = ""
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Fprintf(w, "ParseForm() err: %v", err)
// 			return
// 		}

// 		in := r.FormValue("input")
// 		f := r.FormValue("font")
// 		f += ".txt"

// 		h.Input = in
// 		h.Font = f

// 		result, err := ascii.AsciiDrawer(in, f)
// 		if err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}

// 		h.Output = result
// 		t.ExecuteTemplate(w, "index.html", h)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return
// 	}

// 	// // fmt.Fprintf(w, "%s", Output)
// }

// func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
// 	w.WriteHeader(status)
// 	if status == http.StatusNotFound {
// 		fmt.Fprint(w, "404: Not Found")
// 	}
// }
