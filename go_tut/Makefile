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
ifdef r 
	@echo "...Testing $(p) $(r)"
	go test -v ./$(p) -run $(r)
else ifdef b
	@echo "...Benchmarking $(p)"
	go test -v ./$(p) -bench .
else
	@echo "...Testing $(p)"
	go test -v ./$(p) 
endif
else
	@echo "...Testing"
	go test -v .
endif
