package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type cmdResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is undefined")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homepage)
	mux.HandleFunc("GET /api/v1/date", getDate)

	server := http.Server{Addr: ":" + port, Handler: mux}
	fmt.Printf("Serving on http port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Homie api server"))
}

func getDate(w http.ResponseWriter, r *http.Request) {
	result := cmdResult{}

	out, err := exec.Command("date").Output()
	if err == nil {
		result.Success = true
		result.Message = "The date is " + string(out)
	}

	json.NewEncoder(w).Encode(result)
}
