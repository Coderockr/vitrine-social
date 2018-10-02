.PHONY: all
all: help

# set default as dev if not set
export VITRINESOCIAL_ENV ?= dev
export DATABASE_HOST ?= 0.0.0.0
export m ?= default
export commit ?= HEAD
export bin ?= vitrine-social
export testWatchPort=8091
export GO111MODULE=on

.PHONY: build

setup: ## initial project setup
	cp ./server/config/dev.env.dist ./server/config/dev.env
	go get github.com/rubenv/sql-migrate/...
	make install

org-reset-password: ## reset a organization's password by email
	docker-compose up -d postgres
	cd server && go run main.go org reset-password $(email) $(password)

org-reset-password-on-docker: ## reset a organization's password by email on docker
	docker-compose up -d
	docker-compose exec golang sh -c "cd server && go run main.go org reset-password $(email) $(password)"

update-dev-dependences: # update dev dependences to the most recent
	go get -u github.com/haya14busa/goverage
	go get -u golang.org/x/lint/golint

install: ## install project dependences
	go get github.com/haya14busa/goverage
	go get golang.org/x/lint/golint
	cd server; go mod tidy ; go mod download

install-frontend: ## install frontend dependences
	cd frontend && yarn install

build: install ## builds the application to the paramters bin (bin=vitrine-social)
	cd server && go build -v -o $(bin) .

build-frontend: ## builds frontend application
	cd frontend && yarn install
	cd frontend && yarn build

new-migration: ## create a new migration, use make new-migration m=message to set the message
	sql-migrate new -config=./devops/dbconfig.yml -env=production "$(m)"

migrations: ## run pending migrations
	docker-compose up -d postgres
	go get github.com/rubenv/sql-migrate/...
	sql-migrate up -config=devops/dbconfig.yml -env=$$VITRINESOCIAL_ENV

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
	docker-compose up -d postgres
	go get -u github.com/codegangsta/gin
	cd server; API_PORT=8001 gin --port 8000 --appPort 8001 --bin $(bin) run serve

postgres-cmd: ## open the postgresql command line
	docker-compose exec postgres psql -h $$DATABASE_HOST -U postgres vitrine

postgres-dump: ## dump database
	docker-compose exec postgres sh -c "pg_dump -h $$DATABASE_HOST -U postgres -Fc -f vitrine/.data/vitrine.dump vitrine"

postgres-restore: ## restore database
	docker-compose exec postgres sh -c "pg_restore -h $$DATABASE_HOST -U postgres -v -c --if-exists -d vitrine vitrine/.data/vitrine.dump"

docs-serve: ## start a server with the docs
	cd docs && make serve

docs-build: ## build the docs
	cd docs && make build

docs-open:
	$$BROWSER docs/index.html

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## show source lint
	golint `find server -maxdepth 1 -type d`

tests: ## run go tests
	cd server && go test -v -race ./...

tests-watch:
	go get github.com/smartystreets/goconvey
	cd server && goconvey -port $(testWatchPort)

tests-frontend: ## run frontend tests
	cd frontend && yarn test

coverage: ## outputs coverage to coverage.out
	cd server && goverage -v -race -coverprofile=coverage.out ./...

send-statiscs: ## send statistics to code quality services
	cd server && bash -c "$$(curl -s https://codecov.io/bash)"
	go get -u github.com/schrej/godacov
	cd server && godacov -t ${CODACY_TOKEN} -r ./coverage.out -c $(commit)
