TOPDIR := $(shell pwd)
OUTPUT_DIR := $(TOPDIR)/out/

SOURCES := $(wildcard *.go)
BINARIES := $(patsubst %.go, out/%, $(SOURCES))

all: $(BINARIES)

out/%: %.go
	go build -o $(OUTPUT_DIR) $<

clean:
	rm -rf $(OUTPUT_DIR)
.PHONY: clean
