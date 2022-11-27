GOBIN := $(shell pwd)/bin
PSQL_USER := root
oSQL_PASSWORD := root 
PSQL_DB := cafe
DEV_DSN := postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost/$(POSTGRES_DB)?connect_timeout=10&sslmode=disable
DEV_PSQL_CONTAINER := cafe_db

.PHONY: bin-deps
bin-deps:
	mkdir -p $(GOBIN)/bin
	GOBIN=$(GOBIN) go install github.com/pressly/goose/v3/cmd/goose@v3.7.0
	GOBIN=$(GOBIN) go install github.com/99designs/gqlgen@v0.17.20

.PHONY: generate
generate: 
	./bin/gqlgen generate

.PHONY: run
run:
	go run ./cmd/cafebackend/main.go

.PHONY: new-migration
new-migration:
	./bin/goose -s -dir ./migrations create migration sql
	
.PHONY: dev-db
dev-db: migrate-dev
	sudo docker run -d \
	--name $(DEV_PSQL_CONTAINER) \
	-e POSTGRES_USER=$(POSTGRES_USER) \
	-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	-e POSTGRES_DB=$(POSTGRES_DB) \
	-p 5432:5432 \
	postgres:latest
	
.PHONY: dev-db-down
dev-db-down:
	sudo docker stop cvault_test_db && \
	sudo docker rm cvault_test_db
	
.PHONY: migrate-dev
migrate-dev:
	./bin/goose -dir ./migrations postgres "$(TEST_DSN)" up
