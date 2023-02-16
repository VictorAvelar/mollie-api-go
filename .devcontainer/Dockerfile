ARG GO_VERSION=1-bullseye

ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${GO_VERSION}

RUN apt-get update && \
    apt install -y exa bat curl fontconfig

RUN sh -c "$(curl -fsSL https://starship.rs/install.sh)" -y -f

RUN mkdir /root/.fonts && \
    wget https://github.com/ryanoasis/nerd-fonts/releases/download/v2.3.3/FiraCode.zip && \
    unzip FiraCode.zip && \
    rm FiraCode.zip && \
    fc-cache /root/.fonts

RUN go install github.com/robertkrimen/godocdown/godocdown@latest
