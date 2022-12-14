SHELL = /bin/bash
docker_run_goci = docker run --rm --platform linux/amd64 --name goci gocitools:latest

build:
	docker build --platform linux/amd64 --tag gocitools:latest --file alpine.Dockerfile .

lint:
	$(docker_run_goci) golangci-lint

sec:
	$(docker_run_goci) gosec

nancy:
	$(docker_run_goci) nancy

analyze_gosec:
	$(docker_run_goci) analyze_nancy

analyze_nancy:
	$(docker_run_goci) analyze_nancy
