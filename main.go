package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type UserDetails struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func extractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}

func extractPathParam(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

func greetPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("---- Incoming POST Request ----")
	log.Printf("%s %s", r.Method, r.URL.Path)

	from := extractPathParam(r.URL.Path, "/greet/")
	query := r.URL.Query()
	token := extractBearerToken(r)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		log.Println("Failed to read body:", err)
		return
	}
	defer r.Body.Close()

	log.Printf("Request Body: %s", string(bodyBytes))
	log.Printf("Query Params: %v", query)
	log.Printf("Path Param 'from': %s", from)
	log.Printf("Bearer Token: %s", token)

	var user UserDetails
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println("Invalid JSON:", err)
		return
	}

	response := fmt.Sprintf(
		"Hello, %s! I heard you like %s. From: %s. Query: %v. Token: %s",
		user.Name, user.Hobby, from, query, token)

	log.Printf("Response: %s", response)
	w.Write([]byte(response))
}

func greetGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("---- Incoming GET Request ----")
	log.Printf("%s %s", r.Method, r.URL.Path)

	from := extractPathParam(r.URL.Path, "/greet/")
	query := r.URL.Query()
	token := extractBearerToken(r)

	log.Printf("Query Params: %v", query)
	log.Printf("Path Param 'from': %s", from)
	log.Printf("Bearer Token: %s", token)

	response := fmt.Sprintf(
		"Hello from GET! From: %s. Query: %v. Token: %s",
		from, query, token)

	log.Printf("Response: %s", response)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			greetPostHandler(w, r)
		case http.MethodGet:
			greetGetHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("✅ Server listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
