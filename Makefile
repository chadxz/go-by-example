TOPDIR := $(shell pwd)
OUTPUT_DIR := $(TOPDIR)/out/

APP_SOURCES := $(filter-out $(wildcard *_test.go), $(wildcard *.go))
BINARIES := $(patsubst %.go, out/%, $(APPS_ONLY))

all: $(BINARIES)

out/%: %.go
	go build -o $(OUTPUT_DIR) $<

clean:
	rm -rf $(OUTPUT_DIR)
.PHONY: clean
