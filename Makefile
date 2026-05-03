all: CGTerm

CGTerm: main.go
	go build .
	

run: CGTerm
	./CGTerm

clean:
	rm -f ./CGTerm
