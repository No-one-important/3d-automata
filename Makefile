all:
	go build -o automata $(shell find . -name "*.go")
