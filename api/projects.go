package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShoreLab/shorelab-backend/lib/gateway"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query().Get("fileName")
	g, err := gateway.NewGateway()
	if err != nil {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	s, err := g.Service.GetProjectsService()
	if err != nil {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	json.NewEncoder(w).Encode(s)
}
