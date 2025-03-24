package server

import (
	"encoding/json"
	"net/http"

	storage "github.com/jacobmontes14/montyd/internal/datastore"
)

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
		default:
			http.Error(w, "Method not valid", http.StatusMethodNotAllowed)
		}
	}
}

func (server *Server) addToData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		server.data_store.AddKeyValue(1, "test")
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": "Data was added!"}
		json.NewEncoder(w).Encode(response)

	}
}

func (server *Server) listData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(server.data_store.GetAllKeys())
	}

}
