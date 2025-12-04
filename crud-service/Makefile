

# Run commands
dev:
	go run cmd/main.go
build:
	go build -o bin cmd/main.go
prod:
	bin/main


	

db-test-run:
	docker run --name hackathon-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123321 -e POSTGRES_DB=hackathon -p 5432:5432 -d postgres:15-alpine
db-test-stop:
	docker stop hackathon-postgres
	docker rm hackathon-postgres


swagger:
	swag init --dir cmd,./api,./domain --output ./api/docs
