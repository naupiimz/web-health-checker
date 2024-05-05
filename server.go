package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Address string
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		domain := Domain{}
		err := json.NewDecoder(r.Body).Decode(&domain)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if domain.Port == "" {
			domain.Port = "80"
		}
		io.WriteString(w, domain.Check())
	}
	return
}

func (s *Server) Start() {
	http.HandleFunc("/check", checkHandler)
	log.Fatal(http.ListenAndServe(s.Address, nil))
}
