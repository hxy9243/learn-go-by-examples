GENTOOL := openapi-generator-cli
GIT_HOST := github.com
GIT_USER := hxy9243
GIT_REPO := go-example
GO_SOURCE_DIR := openapi

all: build/server build/client

.PHONY:
gen: gen-server gen-client

.PHONY:
gen-server: api/*.yaml
	$(GENTOOL) generate -g go-server \
		-i api/api.yaml -o server/gen/ \
		--git-host $(GIT_HOST) --git-user-id $(GIT_USER) --git-repo-id $(GIT_REPO) \
		--additional-properties=outputAsLibrary=true,sourceFolder=$(GO_SOURCE_DIR)

	for f in server/gen/openapi/*.go; do \
		echo "formatting $${f}"; \
	    goimports -w $${f}; \
	done

.PHONY:
gen-client: api/*.yaml
	$(GENTOOL) generate -g go \
		-i api/api.yaml -o client/gen/openapi \
		--git-host $(GIT_HOST) --git-user-id $(GIT_USER) --git-repo-id $(GIT_REPO) \
		--additional-properties=outputAsLibrary=true,sourceFolder=$(GO_SOURCE_DIR)

	for f in server/gen/openapi/*.go client/gen/openapi/*.go; do \
		echo "formatting $${f}"; \
	    goimports -w $${f}; \
	done
	rm client/gen/openapi/go.mod client/gen/openapi/go.sum

.PHONY:
doc: api/*.yaml
	$(GENTOOL) generate -g dynamic-html \
		-i api/api.yaml -o server/gen/doc

build/server: server/app/main.go server/service/*.go server/gen/openapi/*.go
	mkdir -p build
	go build -o $@ $<

build/client: client/main.go
	mkdir -p build
	go build -o $@ $<


.PHONY:
clean:
	rm -rf build/