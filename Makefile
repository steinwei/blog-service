dev:
	go run .
build:
	go build . --env prod
prod:
	make build
test:
	go run . --env test