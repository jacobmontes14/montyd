package server

import (
	"encoding/json"
	"net/http"

	storage "github.com/jacobmontes14/montyd/internal/datastore"
)

type Item struct {
	Key   int    `json:"id"`
	Value string `json:"value"`
}

type Server struct {
	server     http.Server
	mux        *http.ServeMux
	data_store *storage.Storage
}

func NewServer(address string) *Server {
	mux := http.NewServeMux()
	server := Server{
		server:     http.Server{Addr: address, Handler: mux},
		mux:        mux,
		data_store: storage.NewDataStore(),
	}

	server.routes()
	return &server
}

func (server *Server) Start() error {
	return server.server.ListenAndServe()
}

func (server *Server) routes() {
	server.mux.HandleFunc("/data", server.handleData())
}

func (server *Server) handleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			server.addToData()(w, r)
		case http.MethodGet:
			server.listData()(w, r)
		case http.MethodDelete:
			server.removeFromData()(w, r)
		default:
			http.Error(w, "Method not valid", http.StatusMethodNotAllowed)
		}
	}
}

func (server *Server) removeFromData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var val Item
		if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		server.data_store.RemoveKey(val.Key)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(val); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (server *Server) addToData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var val Item
		if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		server.data_store.AddKeyValue(val.Key, val.Value)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(val); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (server *Server) listData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(server.data_store.GetAllKeys())
	}
}
