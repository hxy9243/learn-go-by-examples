GENTOOL := openapi-generator-cli
GIT_HOST := github.com
GIT_USER := hxy9243
GIT_REPO := go-example
GO_SOURCE_DIR := openapi

export GO_POST_PROCESS_FILE := "goimports -w"

all: out/server

.PHONY:
gen: api/*.yaml
	$(GENTOOL) generate -g go-server \
		-i api/api.yaml -o server/gen/ \
		--git-host $(GIT_HOST) --git-user-id $(GIT_USER) --git-repo-id $(GIT_REPO) \
		--additional-properties=outputAsLibrary=true,sourceFolder=$(GO_SOURCE_DIR)

	for f in server/gen/openapi/*.go; do \
		echo "formatting $${f}"; \
	    goimports -w $${f}; \
	done

.PHONY:
doc: api/*.yaml
	$(GENTOOL) generate -g dynamic-html \
		-i api/api.yaml -o server/gen/doc

.PHONY:
out/server: server/app/main.go server/service/*.go server/gen/openapi/*.go
	mkdir -p out
	go build -o $@ $<