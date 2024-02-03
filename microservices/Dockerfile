FROM golang:1.20 as dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY greet  ./greet
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./greet/greet_server/server.go
CMD ["/app/server"]
EXPOSE 50051

# creating a min image
FROM scratch as min
COPY --from=dev /app/server /app/server
CMD ["/app/server"]
EXPOSE 50051

