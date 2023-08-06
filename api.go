package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	listenAddr string
	store Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()

  router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)
	if err := http.ListenAndServe(s.listenAddr, router); err != nil {
		log.Fatal("Listen and Serve: ", err)
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
//	id := mux.Vars(r)["id"]
	fmt.Println(mux.Vars(r))
	return WriteJSON(w, http.StatusOK, &Account{})
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

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
