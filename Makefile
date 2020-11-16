
.PHONY: run
run:
	go run cmd/main.go

.PHONY: docker-build
docker-build:
	docker build -t quay.io/didil/bucket-text-api . 

.PHONY: docker-push
docker-push: 
	docker push quay.io/didil/bucket-text-api