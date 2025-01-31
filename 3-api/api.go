package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
            WriteJSON(w, http.StatusBadGateway,  ApiError{Error: err.Error()})
        }
    }
}

type APIServer struct {
    listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
    }
}

func (s *APIServer) Run() {
    router := mux.NewRouter()

    router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

    router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

    log.Printf("JSON API server running on port: %s", s.listenAddr)
    http.ListenAndServe(s.listenAddr, router)
}

func (s * APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

    switch r.Method {
    case "GET":
        return s.handleGetAccount(w,r)
    case "POST":
        return s.handleGetAccount(w,r)
    case "DELETE":
        return s.handleGetAccount(w,r)
    case "PUT":
        return s.handleGetAccount(w,r)
    case "PATCH":
        return s.handleGetAccount(w,r)
    }
    return fmt.Errorf("HTTP Metod not allowed %s", r.Method)
}

func (s * APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    id := mux.Vars(r)["id"]

    fmt.Println(id)

    account := NewAccount("FelipeFarias")
    return WriteJSON(w, http.StatusOK, account)
}

func (s * APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (s * APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}
