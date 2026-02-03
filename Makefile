.PHONY: build help help-doc link-agent test run clean lint deps install docker-build


build:
	go build -o secondbrainmanager && ./secondbrainmanager $(ARGS)

help:
	@echo "Makefile commands:"
	@echo "  build          - Build and run the secondbrainmanager"
	@echo "  help           - Show this help message"
	@echo "  help-doc       - Show documentation-related commands"
	@echo "  link-agent     - Create symbolic links for agent documentation"
	@echo "  test           - Run tests"
	@echo "  run            - Run the secondbrainmanager"
	@echo "  clean          - Clean up build artifacts"
	@echo "  deps           - Download dependencies"
	@echo "  lint           - Run linter"
	@echo "  docker-build   - Build Docker image"
	@echo "  install        - Install secondbrainmanager to GOPATH/bin"

help-doc:
	open README.md

link-agent:
	ln -s AGENTS.md CLAUDE.md
	ln -s AGENTS.md GEMINI.md

test:
	gotestsum --format testname

run:
	go run .

clean:
	rm -f secondbrainmanager

deps:
	go mod download

lint:
	golangci-lint run

docker-build:
	docker build -t secondbrainmanager .

install:
	go build -o $(HOME)/go/bin/secondbrainmanager .
