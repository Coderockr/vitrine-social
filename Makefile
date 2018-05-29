.PHONY: all
all: help

# set default as dev if not set
export VITRINESOCIAL_ENV ?= dev
export DATABASE_HOST ?= 0.0.0.0
export m ?= default
export commit ?= HEAD

.PHONY: build

install: ## install project dependences
	go get github.com/rubenv/sql-migrate/...
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/haya14busa/goverage
	cd server; dep ensure

new-migration: ## create a new migration, use make new-migration m=message to set the message
	sql-migrate new -config=./devops/dbconfig.yml -env=production "$(m)"

migrations: ## run pending migrations
	docker-compose up -d postgres
	go get github.com/rubenv/sql-migrate/...
	sql-migrate up -config=devops/dbconfig.yml -env=production

migrations-on-docker: ## run migrations inside docker
	docker-compose up -d
	docker-compose exec golang sql-migrate up -config=devops/dbconfig.yml -env=docker

serve: ## start server
	docker-compose up -d postgres
	cd server && go run main.go serve

install-on-docker: ## install dependences from docker
	docker-compose up -d
	docker-compose exec golang make install

serve-on-docker: ## start the server inside docker
	docker-compose up -d
	docker-compose exec golang sh -c "cd server && go run main.go serve"

serve-watch: ## start server with hot reload
	docker-compose up -d
	go get -u github.com/codegangsta/gin
	cd server; API_PORT=8001 gin --port 8000 --appPort 8001 --bin server-cmd run serve

postgres-cmd: ## open the postgresql command line
	docker-compose exec postgres psql -h $$DATABASE_HOST -U postgres vitrine

docs-serve: ## start a server with the docs
	cd docs && make serve

docs-build: ## build the docs
	cd docs && make build

docs-open:
	$$BROWSER docs/index.html

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

tests: ## run go tests
	cd server && go test -v -race ./...

coverage: ## outputs coverage to coverage.out
	cd server && goverage -v -race -coverprofile=coverage.out ./...

send-statiscs: ## send statistics to code quality services
	cd server && bash -c "$$(curl -s https://codecov.io/bash)"
	go get -u github.com/schrej/godacov
	cd server && godacov -t ${CODACY_TOKEN} -r ./coverage.out -c $(commit)
