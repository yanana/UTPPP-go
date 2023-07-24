TEST_PACKAGES := $(shell go list -f '{{ if (gt (len .TestGoFiles) 0) }}{{ .ImportPath }}{{ end }}' ./...)
TARGET_FILES := $(shell find . -type f -name "*.go" -print)

.PHONY: test
test: $(TARGET_FILES)
	@go test $(TEST_PACKAGES)
