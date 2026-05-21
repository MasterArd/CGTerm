TARGET = ./build/cgterm

all: cgterm

CGTerm: main.go
	go build -o $(TARGET) main.go

run:
	go clean -cache
	go run -a main.go

build:
	go clean -cache
	go build -o $(TARGET) main.go
	

clean:
	go clean -cache
	rm -r ./build/

install:
	go build -a -o $(TARGET) main.go
	sudo mv ./build/cgterm /usr/bin/

# Requires upx: https://upx.github.io/
compact:
	upx --best ./build/cgterm

