ARG GO_VERSION=1-1.21-bullseye

ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:${GO_VERSION}

RUN go install golang.org/x/tools/cmd/stringer@latest
RUN go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
