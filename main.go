package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Crear enrutador
	router := mux.NewRouter()

	// Manejador para servir los archivos estáticos
	fs := http.FileServer(http.Dir("static"))

	// Ruta para acceder a los archivos estáticos
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Configurar rutas
	router.HandleFunc("/", handlers.Home)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	// Manejo personalizado para página no encontrada (404)
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	// Configuración del servidor
	port := ":8080"

	// Asignar el enrutador al servidor
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(server.ListenAndServe())
}
