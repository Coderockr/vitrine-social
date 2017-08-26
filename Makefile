.PHONY: all
all: build

# set default as dev if not set
export VITRINESOCIAL_ENV ?= dev
export DATABASE_HOST ?= 127.0.0.1

.PHONY: build

install-db:
	docker-compose up -d postgres
	docker-compose exec postgres psql -h $$DATABASE_HOST -U postgres -c "create database vitrine"
	docker-compose exec postgres psql -h $$DATABASE_HOST -U postgres vitrine -f /vitrine/devops/database.sql

install:
	go get github.com/rubenv/sql-migrate/...
	go get -u github.com/golang/dep/cmd/dep
	cd server; dep ensure

migrations:
	go get github.com/rubenv/sql-migrate/...
	sql-migrate up -config=devops/dbconfig.yml -env=production

serve:
	cd server && go run main.go
