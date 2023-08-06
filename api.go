package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error 
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIserver struct {
	listenAddr string
}

func (s *APIserver) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func NewAPIServer(listenAddr string) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
	}
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	//return nil
	// if r.Method == "GET" {
	// 	return s.handleGetAccount(w, r)
	// }
	// if r.Method == "POST" {
	// 	return s.handleCreateAccount(w, r)
	// }
	// if r.Method == "DELETE" {
	// 	return s.handleDeleteAccount(w, r)
	// }
	switch m :=r.Method; m {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
	//return fmt.Errorf("method not allowed %s", r.Method)

	
}
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// func (s *APIserver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }