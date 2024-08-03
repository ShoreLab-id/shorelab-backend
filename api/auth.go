package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShoreLab/shorelab-backend/lib/dto"
	"github.com/ShoreLab/shorelab-backend/lib/gateway"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	auth := &dto.AuthRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	g, err := gateway.NewGateway()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	res, err := g.Service.LoginService(auth)
	if err == mongo.ErrNoDocuments {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	if err == bcrypt.ErrMismatchedHashAndPassword {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong Email or Password"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
