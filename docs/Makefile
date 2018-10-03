.PHONY: all
all: build

AGLIO := $(shell command -v aglio 2> /dev/null)

install:
	npm install -g aglio

.PHONY: build
build: test-install
	aglio -i api.apib --theme-full-width --no-theme-condense -o index.html

test-install:
ifndef AGLIO
	make install
endif

serve: test-install
	aglio -i api.apib --theme-full-width --no-theme-condense -s
