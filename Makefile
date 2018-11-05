ANTLR := java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool

.PHONY: run
run: format
	go run . example.cls

.PHONY: test
test: format
	go test

.PHONY: build
build: format
	go build

.PHONY: format
format:
	# goimports -w .
	gofmt -w .

.PHONY: generate
generate:
	cd ./parser; \
	$(ANTLR) -Dlanguage=Go -visitor apex.g4

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif

.PHONY: deps
deps: dep
	dep ensure

