package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	// "strings"

	"github.com/TheAlok15/email-verification-system/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)


func VerifyHandler(pool *pgxpool.Pool) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		

		var req models.Input

		err := json.NewDecoder(r.Body).Decode(&req)
		if err !=nil{
			http.Error(w, "JSON Invalid", http.StatusBadRequest)
			return
		}

		// if !ValidateEmail(req.Email){
		// 	http.Error(w, "Invalid email format", http.StatusBadRequest)
		// 	return
		// }

		_, err = mail.ParseAddress(req.Email)
		if err != nil {
			http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
		}

		var jobID string

		err = pool.QueryRow(
			r.Context(),
			`INSERT INTO jobs(email, status)
			VALUES($1, 'pending') RETURNING ID`,
			req.Email,


		).Scan(&jobID)

		if err != nil {
			log.Println("db error :", err)
			http.Error(w, "DB insert fails", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"job_id": jobID,
			"status": "pending",
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(response)


	}



}

// func ValidateEmail(email string) bool {

// 	if email == "" {

// 		return false
// 	}

// 	if !strings.Contains(email, "@") {

// 		return false
// 	}

// 	parts := strings.Split(email, "@")

// 	if len(parts) != 2 {
// 		return false
// 	}

// 	if parts[0] == "" || parts[1] == "" {
// 		return false
// 	}

// 	return true



// }