FROM golang:1.16-alpine

WORKDIR /app

COPY . .

EXPOSE 8080
EXPOSE 8081

CMD ["go", "run", "main.go"]
