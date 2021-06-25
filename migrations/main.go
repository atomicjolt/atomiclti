package main

import (
	"log"
	"os"

	"github.com/atomicjolt/atomiclti/repo"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

const directory = "migrations"

func main() {
	db := repo.GetConnection()

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
