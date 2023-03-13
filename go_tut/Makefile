
build:
	go build .

run:
	go run .

test:
ifdef package
	go test -v ./$(package)
else
	go test -v .
endif
