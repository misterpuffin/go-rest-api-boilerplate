NAME=go-rest-api-boilerplate
VERSION=0.0.1

.PHONY: migrate_up
migrate_up: 
	cd sql/migrations; \
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://username:password@localhost:5432/kroshy goose up
	
.PHONY: migrate_down
migrate_down: 
	cd sql/migrations; \
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://username:password@localhost:5432/kroshy goose down


.PHONY: build
build:
	@echo Building from source....

@go build -o ./build/$(NAME)
.PHONY: run
run: build
	@echo Starting your app using dev configs....
	@./build/$(NAME) -e dev

.PHONY: build-prod
build-prod:
	@echo Building from source....
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

.PHONY: run-prod
run-prod:
	@echo Starting app using prod configs....
	@$CI_PROJECT_DIR/go-rest-api-boilerplate-binary -e dev

.PHONY: clean
clean:
	@echo Removing build file....
	@rm -f ./build/$(NAME)

.PHONY: test
	@go test -v ./internal/tests/*
test:
