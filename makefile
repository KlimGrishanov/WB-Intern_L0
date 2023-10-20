db-create:
	docker run --name wb_l0 -e POSTGRES_PASSWORD=passwd -p 5432:5432 -d postgres

migrations-up:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

migrations-down:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' down

run-server:
	go run cmd/main.go