FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY greet  ./greet
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./greet/greet_server/server.go
CMD ["/app/server"]
EXPOSE 50051
