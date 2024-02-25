# Migrate for GitHub Actions

A wrapper for [golang-migrate](https://github.com/golang-migrate/migrate) cli.

## Input variables

See [action.yml](./action.yml) for more detailed information.

* `path` - relative path to your migration folder
* `database` - a command to connect to the database
* `command` - migrate cli command to run
* `prefetch` - Number of migrations to load in advance before executing (default 10)
* `lockTimeout` - Allow N seconds to acquire database lock (default 15)
* `verbose` - Print verbose logging
* `version` - Print version

[Migrate CLI docs](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

## Usage with postgres

```yaml
name: run migration
on: [ push ]
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: vovavc/migrate-github-action@v0.2.1
        with:
          path: ./backend/migrate
          database: postgres://${{ secrets.DB_USER }}:${{ secrets.DB_PASS }}@${{ secrets.DB_HOST }}/${{ secrets.DB_NAME }}?sslmode=disable
          command: up
```

## Usage with mysql

```yaml
name: run migration
on: [ push ]
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: vovavc/migrate-github-action@v0.2.1
        with:
          path: ./backend/migrate
          database: mysql://${{ secrets.DB_USER }}:${{ secrets.DB_PASS }}@tcp(${{ secrets.DB_HOST }})/${{ secrets.DB_NAME }}?tls=${{ secrets.DB_TLS }}
          command: up
```

## Usage with sqlite3

```yaml
name: run migration
on: [ push ]
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: mrmelon54/migrate-action@v0.2.1
        with:
          path: ./backend/migrate
          database: sqlite3://database.sqlite3.db
          command: up
```

## Contributing

Pull requests are welcome at [mrmelon54/migrate-action](https://github.com/mrmelon54/migrate-action/pulls)

## License

The scripts and documentation in this project are released under the [MIT License](LICENSE)
