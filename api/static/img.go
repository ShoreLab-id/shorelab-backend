package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	img, err := http.Get(fmt.Sprintf("https://drive.usercontent.google.com/download?id=%s&export=download&authuser=0", q))
	if img.StatusCode != http.StatusOK {
		w.WriteHeader(img.StatusCode)
		w.Write([]byte(http.StatusText(img.StatusCode)))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	c, err := io.ReadAll(img.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	defer func() {
		err := img.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
	}()

	w.Header().Add("Content-Type", img.Header.Get("Content-Type"))
	w.Write([]byte(c))
}
