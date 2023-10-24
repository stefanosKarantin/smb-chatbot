FROM golang:1.20 as builder

WORKDIR /go/smb-chatbot

COPY . ./

RUN go test -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o smb-chatbot ./cmd/main.go

FROM alpine:3.15
COPY --from=builder /go/smb-chatbot .

ENTRYPOINT ["./smb-chatbot"]