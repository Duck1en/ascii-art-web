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
