FROM golang:1.20.4 as builder

WORKDIR /
COPY go.mod go.sum /
COPY *.go ./
COPY sqlite-migrate/*.go ./sqlite-migrate/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app
RUN CGO_ENABLED=1 GOOS=linux go build -o /migrate-sqlite ./sqlite-migrate/main.go

FROM migrate/migrate

COPY --from=builder ./app ./
COPY --from=builder ./migrate-sqlite ./

ENTRYPOINT ["/app"]
