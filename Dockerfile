FROM golang:1.24.5-alpine

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "-v", "./mollie/...", "-coverprofile", "cover.out"]
