package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheAlok15/email-verification-system/internal/database"
	"github.com/TheAlok15/email-verification-system/internal/models"
)

func main() {

	conn, err := database.Connect()
	if err != nil {
		log.Fatal("Error connection to db ", err)

	}

	defer conn.Close(context.Background())

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.Input
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var jobID string
		err = conn.QueryRow(
			context.Background(),
			`INSERT INTO jobs(email,status)
			VALUES($1, 'pending') RETURNING ID`,
			req.Email,
		).Scan(&jobID)

		if err != nil {
			log.Println("DB ERROR:", err)
			http.Error(w, "DB insert fails", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"job_id": jobID,
			"status": "pending",
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(response)

	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)

}
