package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/search", searchHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter is missing", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("python3", "chatgpt_query.py", query)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
