package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserDetails struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("---- Incoming Request ----")
	log.Printf("%s %s", r.Method, r.URL.Path)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		log.Println("Failed to read body:", err)
		return
	}
	defer r.Body.Close()

	log.Printf("Request Body: %s", string(bodyBytes))

	var user UserDetails
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println("Invalid JSON:", err)
		return
	}

	response := fmt.Sprintf("Hello, %s! I heard you like %s.", user.Name, user.Hobby)
	log.Printf("Response: %s", response)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	log.Println("✅ Server listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
