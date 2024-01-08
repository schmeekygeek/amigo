BINARY=amigo
.DEFAULT_GOAL := run

run:
	go build && ./$(BINARY)
