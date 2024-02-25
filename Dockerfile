FROM golang:1.20.4 as builder

WORKDIR /
COPY go.mod /
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app
RUN CGO_ENABLED=0 GOOS=linux go build ./sqlite-migrate -o /sqlite-migrate

FROM migrate/migrate

COPY --from=builder ./app ./

ENTRYPOINT ["/app"]
