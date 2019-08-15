.PHONY: all
all: help

# set default as dev if not set
export VITRINESOCIAL_ENV ?= dev
export DATABASE_HOST ?= 0.0.0.0
export m ?= default
export commit ?= HEAD
export bin ?= vitrine-social
export testWatchPort=8091

.PHONY: build

go-get-sql-migrate:
	cd server && GO111MODULE=off go get -v github.com/rubenv/sql-migrate/...

./server/config/dev.env.dist:
	cp ./server/config/dev.env.dist ./server/config/dev.env

setup: ./server/config/dev.env.dist install-hostnames go-get-sql-migrate install ## initial project setup

org-reset-password: ## reset a organization's password by email (email=email@email.net password=secret-password)
	docker-compose up -d postgres
	cd server && go run main.go org reset-password $(email) $(password)

org-reset-password-on-docker: ## reset a organization's password by email on docker (email=email@email.net password=secret-password)
	docker-compose up -d
	docker-compose exec golang sh -c "cd server && go run main.go org reset-password $(email) $(password)"

update-dev-dependences: # update dev dependences to the most recent
	cd server && go get -u github.com/haya14busa/goverage
	cd server && go get -u golang.org/x/lint/golint

install: ## install project dependences
	cd server && go get github.com/haya14busa/goverage
	cd server && go get golang.org/x/lint/golint
	make go-get-sql-migrate
	cd server && go mod tidy && go mod download

install-frontend: ## install frontend dependences
	cd frontend && yarn install

build: install ## builds the application to the parameters bin (bin=vitrine-social)
	cd server && go build -v -o $(bin) .

build-frontend: ## builds frontend application
	cd frontend && yarn install
	cd frontend && yarn build

new-migration: ## create a new migration, use make new-migration m=message to set the message
	sql-migrate new -config=./devops/dbconfig.yml -env=production "$(m)"

migrations: ## run pending migrations
	docker-compose up -d postgres
	make go-get-sql-migrate
	sql-migrate up -config=devops/dbconfig.yml -env=$$VITRINESOCIAL_ENV

migrations-on-docker: ## run migrations inside docker
	docker-compose up -d
	docker-compose exec golang sql-migrate up -config=devops/dbconfig.yml -env=docker

start-dependences: ## docker up all container dependencies
	docker-compose up -d postgres minio images-server

serve: start-dependences ## start server
	cd server && go run main.go serve

install-hostnames:
	echo -e '127.0.0.1 api.vitrinesocial.test # usar porta 8000 (golang)' | sudo tee -a /etc/hosts
	echo -e '127.0.0.1 images.vitrinesocial.test # usar porta 7000 (images-server)' | sudo tee -a /etc/hosts
	echo -e '127.0.0.1 minio.vitrinesocial.test # usar porta 9000 (minio)' | sudo tee -a /etc/hosts
	echo -e '127.0.0.1 vitrinesocial.test # usar porta 3000 (frontend)' | sudo tee -a /etc/hosts

setup-on-docker: install-hostnames install-on-docker ## setups the backend project

install-on-docker: ./server/config/dev.env.dist ## install dependences from docker
	docker-compose up -d
	docker-compose exec golang make install

serve-on-docker: ## start the server inside docker
	docker-compose up -d
	docker-compose exec golang sh -c "cd server && go run main.go serve"

serve-watch: start-dependences ## start server with hot reload (bin=vitrine-social)
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

docs-open: ## opens the docs on your browser
	$$BROWSER docs/index.html

open-minio: ## opens the docs on your browser
	$$BROWSER localhost:9000

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
