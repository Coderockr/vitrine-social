.PHONY: all
all: build

.PHONY: build

install-db:
	docker run -it --rm --link postgres:postgres postgres:9-alpine psql -h postgres -U postgres -c "create database vitrine"
	docker run -it --rm --link postgres:postgres -v $(pwd):/vitrine postgres:9-alpine psql -h postgres -U postgres vitrine -f /vitrine/devops/database.sql
