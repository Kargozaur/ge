include makefiles/migrations.mk

.PHONY: run build fmt

run:
	go run cmd/main.go

build:
	go build -o app cmd/main.go

fmt:
	go fmt ./..