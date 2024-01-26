.PHONY: build
build:
	@echo "Building..."
	go build 

.PHONY: run
run:
	@echo "Running..."
	go run main.go

.PHONY: test
test:
ifdef t 
	@echo "...Testing $(t)"
	go test -v ./$(t)
else
	@echo "...Testing"
	go test -v .
endif
