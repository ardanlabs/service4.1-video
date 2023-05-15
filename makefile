# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run-local:
	go run app/services/sales-api/main.go

tidy:
	go mod tidy
	go mod vendor