# Makefile

# Define targets
lint:
	golint ./..
test:
	go test -v
build:
	go build -o myapp ./main.go
clean:
	rm -f myapp