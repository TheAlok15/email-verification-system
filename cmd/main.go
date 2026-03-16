package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TheAlok15/email-verification-system/internal/database"
	"github.com/TheAlok15/email-verification-system/internal/handler"
	"github.com/TheAlok15/email-verification-system/internal/worker"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	fmt.Println(os.Getenv("DB_URL"))

	pool, err := database.Connect()
	if err != nil {
		log.Fatal("Error connection to db ", err)

	}

	defer pool.Close()

	ctx := context.Background()

	w := worker.NewWorker(pool)

	w.Start(ctx, 5)

	http.HandleFunc("/verify", handler.VerifyHandler(pool))

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)

}
