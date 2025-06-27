MODULE_DIRS := $(shell find . -name "go.mod" -exec dirname {} \;)

.PHONY: deps

deps:
	@for DIR in $(MODULE_DIRS); do \
		cd $$DIR && go get ./... && go mod tidy && go mod verify; \
	done