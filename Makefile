all: CGTerm

CGTerm: main.go
	go build .
	

run: CGTerm
	clear && ./CGTerm

clean:
	rm -f ./CGTerm
