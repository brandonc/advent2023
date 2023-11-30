build: generate
	go build -o bin/advent2023 cmd/main.go

generate:
	go generate internal/commands/init.go

today:
	ARG=today go generate internal/commands/init.go

tomorrow:
	ARG=tomorrow go generate internal/commands/init.go

.PHONY: build generate today tomorrow
