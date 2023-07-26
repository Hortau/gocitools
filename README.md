# gocitools

You can run test and scan for vulnerability.
- [Go](https://go.dev/) 1.x Alpine Linux
- [gosec](https://github.com/securego/gosec) - Golang Security Checker
- [Nancy](https://github.com/sonatype-nexus-community/nancy) - check Golang dependencies
- analyze_gosec script to use in CI
- analyze_nancy script to use in CI

## Build the container image
    make build

## Run golangci-lint
    make lint

## Run analyze gosec script
    make analyze_gosec

## Run analyze nancy script
    make analyze_nancy

## Remove the container image
    make clean