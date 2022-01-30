run:
	go run cmd/app/main.go

test:
	go test ./...

tidy:
	go mod tidy

create-docker-image:
	docker build -t balrammani/url_shortner .

docker-run:
	docker run --rm -p 8080:8080 --name short-url rest-url-shortner

docker-clean:
	docker stop short-url