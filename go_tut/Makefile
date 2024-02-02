.PHONY: build
build:
	@echo "Building..."
	go build 

.PHONY: run
run:
	@echo "Running..."
	go run main.go

.PHONY: check
check:
	@echo "Checking..."
	go build -o /dev/null
	go vet

.PHONY: test
test:
ifdef p 
ifdef t 
	@echo "...Testing $(p) $(t)"
	go test -v ./$(p) -run $(t)
else
	@echo "...Testing $(t)"
	go test -v ./$(p) 
endif
else
	@echo "...Testing"
	go test -v .
endif
