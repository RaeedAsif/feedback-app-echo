run:
	go run main.go
migrate-db:
	go run migrate/migrate.go
swag:
	swag init main.go --output docs/