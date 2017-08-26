.PHONY: all
all: build


.PHONY: build
build: 
	aglio -i api.apib --theme-full-width --no-theme-condense -o index.html

new:
	mkdir blueprint/$(filter-out $@,$(MAKECMDGOALS))
	> blueprint/$(filter-out $@,$(MAKECMDGOALS))/api.apib
	> blueprint/$(filter-out $@,$(MAKECMDGOALS))/data-structures.apib

%:      # ignora o segundo parâmetro do new
	@:

install:
	npm install -g aglio

serve:
	aglio -i api.apib --theme-full-width --no-theme-condense -s
