default:
	cat Makefile

run:
	go run cmd/*

test:
	go test ./...