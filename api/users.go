package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/ShoreLab/shorelab-backend/lib/dto"
	"github.com/ShoreLab/shorelab-backend/lib/gateway"
)

var ErrUserExists = errors.New("user already exists")

func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if r.Method != http.MethodPost {
			log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
			return
		}
	}

	g, err := gateway.NewGateway()
	if err != nil {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	u := &dto.UserCreateRequest{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err = g.Service.CreateUserService(u)
	if err == ErrUserExists {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	log.Default().Println(r.RemoteAddr, r.RequestURI, http.StatusCreated, http.StatusText(http.StatusCreated))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created"))
}
