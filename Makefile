GOBIN := $(shell pwd)/bin
PSQL_USER := root
PSQL_PASSWORD := root
PSQL_DB := cafe
DEV_DSN := postgresql://$(PSQL_USER):$(PSQL_PASSWORD)@localhost/$(PSQL_DB)?connect_timeout=10&sslmode=disable
CONTAINER_NAME := cafe_db

.PHONY: bin-deps
bin-deps:
	mkdir -p $(GOBIN)/bin
	GOBIN=$(GOBIN) go install github.com/pressly/goose/v3/cmd/goose@v3.7.0
	GOBIN=$(GOBIN) go install github.com/99designs/gqlgen@v0.17.20

.PHONY: generate
generate: bin-deps
	./bin/gqlgen generate

.PHONY: run
run:
	go run ./cmd/cafebackend/main.go

.PHONY: new-migration
new-migration:
	./bin/goose -s -dir ./migrations create migration sql
	
.PHONY: dev-db
dev-db: 
	docker run -d \
	--name $(CONTAINER_NAME) \
	-e POSTGRES_USER=$(PSQL_USER) \
	-e POSTGRES_PASSWORD=$(PSQL_PASSWORD) \
	-e POSTGRES_DB=$(PSQL_DB) \
	-p 5432:5432 \
	postgres:latest
	
.PHONY: dev-db-down
dev-db-down:
	docker stop $(CONTAINER_NAME) && \
	docker rm $(CONTAINER_NAME) 
	
.PHONY: migrate-dev
migrate-dev:
	./bin/goose -dir ./migrations postgres "$(DEV_DSN)" up
