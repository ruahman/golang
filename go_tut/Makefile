.PHONY: build
build:
	@echo "Building..."
	go build 

.PHONY: run
run:
	echo "Running..."
	go run main.go

.PHONY: test
test:
ifdef test 
	@echo "...Testing $(test)"
	go test -v ./$(test)
else
	@echo "...Testing"
	go test -v .
endif
