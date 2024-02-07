FROM golang:1.22.0-alpine3.19 AS dev

ENV APP_VERSION ${APP_VERSION:-dev}
ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY api api
COPY database database
COPY internal internal
COPY main.go main.go

ENTRYPOINT [ "go", "run", "main.go" ]

CMD [ "serve" ]

FROM dev AS builder

RUN go build -o bin/todod main.go

ENTRYPOINT []

FROM scratch AS bin

COPY --from=builder /app/bin/todod /todod

ENTRYPOINT [ "/todod" ]

CMD [ "serve" ]
