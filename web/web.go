package web

import (
	"os"

	"abiiram10/HelloHeroku/web/usuarios"

	"log"
	"net/http"
)

func Ini() {
	http.HandleFunc("/obtenerUsuarios", usuarios.GetUsuariosPOST)
	http.HandleFunc("/nuevoUsuario", usuarios.SetUsuarioPOST)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("Server starting at port " + port + "...")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
