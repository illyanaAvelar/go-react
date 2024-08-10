package gen

//go:generate go run ./cmd/tools/terndotenv/main.go
//go:generate C:/Users/illya/go/bin/sqlc.exe generate -f ./internal/store/pgstore/sqlc.yaml