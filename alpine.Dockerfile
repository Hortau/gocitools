FROM golang:1.19-alpine

ARG lint_version=v1.50.1
ARG gosec_version=v2.14.0
ARG nancy_version=v1.0.24

ENV GOROOT /usr/local/go

RUN apk --no-cache add git

RUN wget -q -O - https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $lint_version
RUN wget -q -O - https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s $gosec_version
RUN wget -q -c https://github.com/sonatype-nexus-community/nancy/releases/download/$nancy_version/nancy-$nancy_version-linux-amd64.tar.gz -O - | tar -xz -C /usr/bin/

WORKDIR /build
COPY /cmd .
RUN cd analyze_gosec && go build -ldflags "-s -w" -o /usr/bin/analyze_gosec main.go
RUN cd analyze_nancy && go build -ldflags "-s -w" -o /usr/bin/analyze_nancy main.go
RUN rm -rf /build

CMD ["golangci-lint", "gosec", "nancy", "analyze_gosec", "analyze_nancy"]
