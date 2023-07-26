# gocitools

You can run test and scan for vulnerability.
- [Go](https://go.dev/) 1.x Alpine Linux
- [gosec](https://github.com/securego/gosec) - Golang Security Checker
- [Nancy](https://github.com/sonatype-nexus-community/nancy) - check Golang dependencies
- analyze_gosec script to use in CI
- analyze_nancy script to use in CI

## How to test with Docker Desktop
Use the commands in the Makefile

`make build, make lint, make sec, make nancy`