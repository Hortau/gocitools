SHELL = /bin/bash

build:
	docker build --platform linux/amd64 --tag gocitools:latest --file alpine.Dockerfile .

lint:
	docker run --platform linux/amd64 --name goci -it gocitools:latest golangci-lint

sec:
	docker run --platform linux/amd64 --name goci -it gocitools:latest gosec

nancy:
	docker run --platform linux/amd64 --name goci -it gocitools:latest nancy

clean:
	docker stop goci
	docker rm goci
