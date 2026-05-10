all: CGTerm

CGTerm: main.go
	go build .

run:
	go clean -cache
	go run -a .

build:
	go clean -cache
	go build -a -o CGTerm .

clean:
	go clean -cache
	rm -f CGTerm
