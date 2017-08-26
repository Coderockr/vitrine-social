.PHONY: all
all: build

.PHONY: build

install-db:
	# docker run -it --rm --link postgres:postgres postgres:9-alpine psql -h postgres -U postgres -c "create database vitrine"
	docker run -it --rm --link postgres:postgres -v ${PWD}:/vitrine postgres:9-alpine psql -h postgres -U postgres vitrine -f /vitrine/devops/database.sql
install: generatekey
	go get github.com/rubenv/sql-migrate/...
	go get -u github.com/golang/dep/cmd/dep
	cd server; dep ensure
migrations:
	sql-migrate up -config=devops/dbconfig.yml -env=production 
serve:
	cd server && go run main.go

generatekey:
	openssl genrsa -out server/config/rsa.key 1024
	openssl rsa -in server/config/rsa.key -pubout > server/config/rsa.key.pub