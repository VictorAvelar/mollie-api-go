ARG GO_VERSION=1-bookworm

ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:${GO_VERSION}

RUN apt-get update && apt-get install -y exa bat curl && apt-get clean

RUN git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting \
    && git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
RUN go install github.com/goreleaser/goreleaser/v2@latest