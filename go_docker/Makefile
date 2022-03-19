.PHONY: image-dev
image-dev:
ifndef rmi
	docker build --target development -t golang-dev .
else	
	docker image rm golang-dev
endif

.PHONY: image-debug
image-debug:
ifndef rmi
	docker build --target debug -t golang-debug .
else
	docker image rm golang-debug
endif

.PHONY: shell
shell:
	docker run -it -v $(shell pwd):/go/src/work --rm golang-dev /bin/bash

.PHONY: debug
debug:
	docker run -it -p 2345:2345 -v $(shell pwd):/go/src/work --rm --security-opt seccomp=unconfined golang-debug /bin/bash