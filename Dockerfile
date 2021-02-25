FROM golang:1.16

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "-v", "./mollie/...", "-coverprofile", "cover.out"]
