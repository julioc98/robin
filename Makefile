test:
	go test ./... -v -coverprofile fmtcoverage.html fmt

test/docker:
	docker-compose run robin make test

run/api:
	go run cmd/api/main.go

run/db:
	docker-compose up robindb

run/adminer:
	docker-compose up adminer

run/docker:
	docker-compose up --build

clean/docker:
	docker-compose down 