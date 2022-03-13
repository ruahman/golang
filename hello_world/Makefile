tag = 1.0.0

.PHONY: image-debug
image-debug:
	docker build --target debug -t golang-debug:$(tag) .

.PHONY: image-dev
image-dev:
	docker build --target dev -t golang-dev:$(tag) .

.PHONY: shell
shell:
	docker run -it -v $(shell pwd):/app --rm golang-dev:$(tag) /bin/bash