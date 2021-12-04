package api

import (
	"cat-bookmarks/domain/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	srv    *http.Server
	bmCase *usecase.BookmarkCase
}

func NewServer(port string, bmc *usecase.BookmarkCase) *Server {
	hs := &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	return &Server{
		srv:    hs,
		bmCase: bmc,
	}

}

func (s *Server) SetRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/bookmarks", s.ShowBookmarks).Methods(http.MethodGet)
	s.srv.Handler = r
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}
