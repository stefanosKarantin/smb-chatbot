package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stefanosKarantin/smb-chatbot/internal/app"
	ihttp "github.com/stefanosKarantin/smb-chatbot/internal/http"
	"github.com/stefanosKarantin/smb-chatbot/internal/storage"
)

func main() {
	storage := storage.NewStorage()

	client := http.Client{}
	host := os.Getenv("MESSAGE_HOST")
	app := app.NewService(&storage.PromotionStorage, &storage.StatsStorage, client, host)

	handler := ihttp.NewHandler(app)
	handler.AppendRoutes()

	err := http.ListenAndServe(":8080", handler.Router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on port 8080")
}
