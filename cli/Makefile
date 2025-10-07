.PHONY: build

build:
	go build -o bin/trustmint main.go


.PHONY: docker-build-all docker-build-base docker-build-sklearn docker-build-pytorch docker-build-tensorflow

docker-build-base:
	docker build -t trustmint/train-base:latest -f docker/train-base/Dockerfile docker/train-base

docker-build-sklearn:
	docker build -t trustmint/train-sklearn:latest -f docker/train-sklearn/Dockerfile docker/train-sklearn

docker-build-pytorch:
	docker build -t trustmint/train-pytorch:latest -f docker/train-pytorch/Dockerfile docker/train-pytorch

docker-build-tensorflow:
	docker build -t trustmint/train-tensorflow:latest -f docker/train-tensorflow/Dockerfile docker/train-tensorflow

docker-build-all: docker-build-base docker-build-sklearn docker-build-pytorch docker-build-tensorflow
