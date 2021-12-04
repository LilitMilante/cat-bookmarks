package api

import (
	"fmt"
	"net/http"
)

func (s *Server) ShowBookmarks(w http.ResponseWriter, _ *http.Request) {
	bookmarks, err := s.bmCase.AllBookmarks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	fmt.Fprint(w, bookmarks)
}
