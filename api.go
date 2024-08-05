// Aqui se inicializa el servidor y todo lo relacionado con eso.
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	//Es el resositorio que almacenara todas las conexiones a la base de datos.
	store *Store
}
//Constructor
func NewAPIServer(addr string, store *Store) *APIServer {
	return &NewAPIServer{addr: addr, store: store}
}
//Methods
//Inicializa el router, registra todos los servicios y escucha al servidor.
func (s *APIServer) Serve() {
	r := mux.NewRouter()

	sb := r.PathPrefix("/api/v1").Subrouter()

	//registering our services

	log.Println("Starting the API server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, sb))
}