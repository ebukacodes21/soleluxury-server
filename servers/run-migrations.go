package servers

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration(dbSource, dbUrl string) {
	migration, err := migrate.New(dbSource, dbUrl)
	if err != nil {
		log.Fatal("unable to create migration", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Unable to run migration:", err)
	}

	log.Print("migration ran successfully")
}
