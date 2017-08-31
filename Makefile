.PHONY: all
all: build

# set default as dev if not set
export VITRINESOCIAL_ENV ?= dev
export DATABASE_HOST ?= 0.0.0.0

.PHONY: build

install:
	go get github.com/rubenv/sql-migrate/...
	go get -u github.com/golang/dep/cmd/dep
	cd server; dep ensure

new-migration:
	sql-migrate new -config=./devops/dbconfig.yml -env=production $(m)

migrations:
	docker-compose up -d
	go get github.com/rubenv/sql-migrate/...
	sql-migrate up -config=devops/dbconfig.yml -env=production

serve:
	docker-compose up -d
	cd server && go run main.go

serve-watch:
	docker-compose up -d
	go get github.com/codegangsta/gin
	gin --port 8081 --appPort 8000 --path ${PWD}/server --bin server-cmd run server/main.go

postgres-cmd:
	docker-compose exec postgres psql -h $$DATABASE_HOST -U postgres vitrine