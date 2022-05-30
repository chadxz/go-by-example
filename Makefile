TOPDIR := $(shell pwd)
OUTPUT_DIR := $(TOPDIR)/out

SOURCES := $(wildcard *.go)
BINARIES := $(patsubst %.go, out/%, $(SOURCES))

all: $(BINARIES)

$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

out/%: %.go $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR) $<

clean:
	rm -rf $(OUTPUT_DIR)
.PHONY: clean
