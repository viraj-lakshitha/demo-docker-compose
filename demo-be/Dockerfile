FROM golang:latest
WORKDIR /app

COPY . .
RUN go get ./...

EXPOSE 8000
ENTRYPOINT ["go","run","main.go"]
