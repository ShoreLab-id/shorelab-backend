package handler

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/ShoreLab/shorelab-backend/lib/gateway"
)

type ErrorMsg []string

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("fileName")
	if q == "" {
		e := map[string]ErrorMsg{
			"errors": {"Missing required query param 'fileName'."},
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	g, err := gateway.NewGateway()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	c, ct, err := g.Service.GetImage(q)
	if err == storage.ErrObjectNotExist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Add("Content-Type", ct)
	w.Write([]byte(c))
}
