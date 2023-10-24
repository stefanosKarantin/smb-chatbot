package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

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

	port := os.Getenv("PORT")
	fmt.Printf("Ready to listen on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler.Router)
	if err != nil {
		log.Fatal(err)
	}
}
