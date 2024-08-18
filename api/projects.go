package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShoreLab/shorelab-backend/lib/gateway"
	"github.com/ShoreLab/shorelab-backend/lib/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	q := r.URL.Query().Get("projectID")
	g, err := gateway.NewGateway()
	if err != nil {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusText(http.StatusInternalServerError))
		return
	}

	if q != "" {
		s, err := g.Service.GetProjectDetailsService(q)
		if err == repository.ErrInvalidID {
			log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		if err == mongo.ErrNoDocuments {
			log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(http.StatusText(http.StatusNotFound))
			return
		}
		if err != nil {
			log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusText(http.StatusInternalServerError))
			return
		}
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusOK, http.StatusText(http.StatusOK))
		json.NewEncoder(w).Encode(s)
		return
	}

	s, err := g.Service.GetProjectsService()
	if err != nil {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusOK, http.StatusText(http.StatusOK))
	json.NewEncoder(w).Encode(s)
}
