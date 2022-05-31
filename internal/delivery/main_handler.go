package delivery

import (
	"net/http"

	"asciiart/internal/ascii"
)

func HomePage(page http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(page, http.StatusMethodNotAllowed, ErrWrongMethod)
		return
	}

	font := r.FormValue("banner")
	text := r.FormValue("input")
	color := r.FormValue("color")
	if text == "" {
		ErrorHandler(page, http.StatusBadRequest, ErrBadRequest)
		return
	}

	finalArt, code, err := ascii.AsciiDrawer(text, font)
	if err != nil {
		ErrorHandler(page, code, err)
		return
	}

	d := struct {
		Data  string
		Color string
	}{
		Data:  finalArt,
		Color: color,
	}

	MainTmpl.Execute(page, d)
}
