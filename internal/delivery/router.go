package delivery

import (
	"net/http"
	"text/template"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Router() *http.ServeMux {
	var err error
	MainTmpl, err = template.ParseFiles("templates/index.html")
	ErrorCheck(err)
	ErrTmpl, err = template.ParseFiles("templates/error.html")
	ErrorCheck(err)
	s.mux.HandleFunc("/", PathPage)
	s.mux.HandleFunc("/ascii-art-web", HomePage)
	s.mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	s.mux.Handle("/templates/images/", http.StripPrefix("/templates/images/", http.FileServer(http.Dir("./templates/images/"))))
	return s.mux
}

func PathPage(page http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(page, http.StatusNotFound, ErrNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(page, http.StatusBadRequest, ErrBadRequest)
		return
	}
	MainTmpl.Execute(page, nil)
}
