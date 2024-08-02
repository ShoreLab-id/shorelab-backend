package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShoreLab/shorelab-backend/lib/gateway"
	"go.mongodb.org/mongo-driver/mongo"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	uname, pw := r.FormValue("username"), r.FormValue("password")
	log.Default().Println(uname, pw)

	g, err := gateway.NewGateway()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	res, err := g.Service.GetUser()
	if err == mongo.ErrNoDocuments {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	json.NewEncoder(w).Encode(res)
}
