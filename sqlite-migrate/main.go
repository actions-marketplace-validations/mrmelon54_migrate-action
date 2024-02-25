package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func main() {
	pathFlag := flag.String("path", "", "")
	databaseFlag := flag.String("database", "", "")
	flag.String("prefetch", "", "")
	flag.String("lock-timeout", "", "")
	flag.String("verbose", "", "")
	flag.String("version", "", "")
	flag.Parse()
	args := flag.Args()

	m, err := migrate.New("file://"+*pathFlag, *databaseFlag)
	if err != nil {
		return
	}

	steps := false
	var n uint64
	if len(args) >= 2 {
		steps = true
		n, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			panic(err)
		}
	}
	switch args[0] {
	case "up":
		if steps {
			err = m.Steps(int(n))
		} else {
			err = m.Up()
		}
	case "down":
		if steps {
			err = m.Steps(-int(n))
		} else {
			err = m.Down()
		}
	default:
		panic("invalid command")
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("Migrate successful")
}
