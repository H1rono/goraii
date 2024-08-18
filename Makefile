all: lib examples

.PHONY: lib
lib:
	go build

file-open: lib examples/file-open/main.go
	go build -o ./build/file-open ./examples/file-open

.PHONY: examples
examples: lib file-open
