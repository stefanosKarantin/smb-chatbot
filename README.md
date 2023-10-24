## SMB-CHATBOT

This coding exersice refers to connectly's assessment for the backend engineer position.

The requirements include building a simple chatbot with a predifined flow.

In this source code we have created a simple http client in order to send messages through an `X` messaging platform. On the other hand we expose some endpoints to have a way of receiving messages and notifications from the platform, like having defined webhooks that refer to these endpoints. 

The implementation was based in extensibility, thats why there are explicit layers on storage, application and http levels. 

The domain is consider to be the promotion and stats entities.

The initialization of each flow is made by an http request on the `/start-promotion` endpoint, like creating a new promotion from an internal tool, just the way a sales/operation person would do.

A Makefile has been added to simplify the run, deployment (if on heroku) and local development:
* test: Runs tests for the project using the Go testing framework.
* run: Runs the application using the go run command.
* fmt: Formats the code using the go fmt command.
* serve: Starts the application using Docker.
* stop: Stops the Docker container running the application.
* heroku: Deploys the application to Heroku (you need to login as well).
* help: Displays a list of available commands and their descriptions.