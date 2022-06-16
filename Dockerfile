FROM golang:1.18

WORKDIR /app/pokedex

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .
RUN go mod download && go mod verify

RUN go build -o ./pokedex-app main.go

EXPOSE 8080