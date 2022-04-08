package main

import (
	"abiiram10/helloheroku/datos"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {

	datos.Ini()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("Server starting at port " + port + "...")

	http.HandleFunc("/hello", helloHeroku)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHeroku(Response http.ResponseWriter, Request *http.Request) {
	defer Request.Body.Close()

	RespuestaBytes, err_marshal := json.Marshal(map[string]interface{}{
		"Resultado": "hello",
	})

	if err_marshal != nil {
		Response.WriteHeader(http.StatusBadRequest)
		return
	}

	Response.Header().Set("Access-Control-Allow-Origin", "*")
	Response.Header().Set("Content-Type", "application/json")
	Response.WriteHeader(http.StatusOK)
	Response.Write(RespuestaBytes)
}
