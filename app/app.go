package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //
	}

	fmt.Println("Servido v2 corriendo en el puerto:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error al arrancar:", err)
		os.Exit(1)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("APP_MESSAGE")
	apiKey := os.Getenv("API_KEY")

	fmt.Fprintf(w, "App de Go v2\n")
	fmt.Fprintf(w, "Mensaje desde ConfigMap: %s\n", message)
	fmt.Fprintf(w, "API Key desde Secret: %s\n", apiKey)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
