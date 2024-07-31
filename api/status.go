package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	res := map[string]string{"status": "OK"}
	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err.Error())
	}
	w.Write(b)
}
