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
	store Store
}

// Constructor
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// Methods
func (ap *APIServer) Serve() {
	//Inicializa el router, registra todos los servicios y escucha al servidor.
	r := mux.NewRouter()

	sb := r.PathPrefix("/api/v1").Subrouter()

	//registering our services
	usersService := NewUserService(ap.store)
	usersService.RegisterRoutes(sb)
	
	tasksService := NewTasksService(ap.store)
	tasksService.RegisterRoutes(sb)


	log.Println("Starting the API server at", ap.addr)

	log.Fatal(http.ListenAndServe(ap.addr, sb))
}
