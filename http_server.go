package workday

import (
	"encoding/json"
	"net/http"
)

type HTTPServer struct {
	store DateStore
	http.Handler
}

func NewHTTPServer() *HTTPServer {
	s := new(HTTPServer)
	s.store = NewMemoryDateStore()

	router := http.NewServeMux()

	router.Handle("/holiday/", http.HandlerFunc(s.dateHandler))

	s.Handler = router
	return s
}

func (s *HTTPServer) dateHandler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Path[len("/holiday/"):]
	dayInfo := s.store.Search(date)
	json.NewEncoder(w).Encode(dayInfo)
	//fmt.Fprint(w, dayInfo)
}
