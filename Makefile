all: CGTerm

CGTerm: main.go
	go build .

run:
	go clean -cache
	go run -a .

build:
	go clean -cache
	go build -a -o cgterm .
	

clean:
	go clean -cache
	rm -f cgterm

install:
	go clean -cache
	go build -a -o cgterm .
	sudo mv cgterm /usr/bin/

compact:
	upx --best ./CGTerm
