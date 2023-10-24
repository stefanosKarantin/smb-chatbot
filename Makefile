test:
	@echo "Running tests..."
	@go test -v ./...

run:
	@echo "Running the application..."
	@go run cmd/main.go

fmt:
	@echo "Formatting the code..."
	@go fmt ./...

serve:
	@echo "Starting the application with Docker..."
	@docker build -t chatbot . && docker run -p 8080:8080 -it chatbot 

stop:
	@echo "Stopping the application with Docker..."
	@docker stop chatbot

heroku: 
	@heroku login
	@heroku create
	@git push heroku main

help:
	@echo "Available commands:"
	@echo "  test   - Run tests"
	@echo "  run    - Run the application"
	@echo "  fmt    - Format the code"
	@echo "  serve  - Start the application with Docker"
	@echo "  stop   - Stop the Docker application"
	@echo "  heroku - Deploy to Heroku"

