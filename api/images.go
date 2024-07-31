package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ShoreLab/shorelab-backend/lib/gateway"
)

func Image(w http.ResponseWriter, r *http.Request) {
	g, err := gateway.NewGateway()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	res, err := g.Service.GetImageList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	json.NewEncoder(w).Encode(res)
}
