package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

// Structura para jugador
type Player struct {
	Name string
}

// Creando jugador
var player Player

// Controlador de inicio
func Home(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "home.html", nil)
}

// Controlador de nuevo juego
func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "new-game.html", nil)
}

// Controlador de juego
func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// Leer los datos del formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	// Redirecíon a otra ruta
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, "game.html", player)
}

// Controlador de jugar
func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Controlador de about
func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}

// Manejo de páginas de error
var errorTemplates = template.Must(template.ParseGlob("templates/**/*.html"))

func handlerError(w http.ResponseWriter, name string, status int) {
	w.WriteHeader(status)
	errorTemplates.ExecuteTemplate(w, name, nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Devolver un error personalizado para páginas no encontradas
	handlerError(w, "404", http.StatusNotFound)
}

// Renderizar plantillas HTML
const baseDir = "templates/"

func renderTemplate(w http.ResponseWriter, name string, data any) {
	templates := template.Must(template.ParseFiles(baseDir+"base.html", baseDir+name))
	// Encabezado
	w.Header().Set("Content-Type", "text/html")

	// Renderizar la plantilla en la respuesta
	err := templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		handlerError(w, "500", http.StatusInternalServerError)
	}
}

// Reiniciar valores
func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
