FROM golang:1.20.4 as builder

WORKDIR /
COPY go.mod /
COPY *.go ./
COPY sqlite-migrate/*.go ./sqlite-migrate
RUN CGO_ENABLED=0 GOOS=linux go build -o /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /sqlite-migrate ./sqlite-migrate

FROM migrate/migrate

COPY --from=builder ./app ./

ENTRYPOINT ["/app"]
