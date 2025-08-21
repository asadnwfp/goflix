# Dockerfile
From golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o movieServer .

EXPOSE 8000

CMD ["./movieServer"]