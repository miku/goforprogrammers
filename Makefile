SHELL := /bin/bash
MD_FILES = $(shell find . -name '*.md')

.PHONY: embedmd
embedmd: $(MD_FILES)
	find . -name "*.md" -exec embedmd -w {} \;

