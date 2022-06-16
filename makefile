run:
	go run main.go "postgres://postgres:1234@localhost:5000/postgres?sslmode=disable"

gen:
	go get github.com/99designs/gqlgen@v0.17.10

	go run github.com/99designs/gqlgen generate

build:
	docker build -t pokedex-app:1.0 .

pokedex-docker:
	docker run -p8080:5432 -e POSTGRES_USERNAME=postgres -e POSTGRES_PASSWORD=1234 --name postgres-pokedex --net postgres-network postgres
	docker run -p8080:8080 --name pokedex-app --net pokedex-network pokedex-app:1.0