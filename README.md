DB_HOST=localhost:3306 DB_USER=root DB_PASSWORD=credential DB_SCHEMA=grab go run server/main.go
SERVER=localhost:50051 go run client/main.go
