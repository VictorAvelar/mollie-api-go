FROM golang:1.18.3-alpine

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "-v", "./mollie/...", "-coverprofile", "cover.out"]
