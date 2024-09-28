.PHONY: all
all:
	go build

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean: tidy
	rm ./bingo-factory

.PHONY: test
test:
	go test ./...
