package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationVersion = 1

func migrateUp(sUrl, dbUrl string) error {
	log.Println("Starting database migrations...")

	// get source url
	log.Println("Migrating source: " + sUrl)

	// get database url
	log.Println("Migrating target: " + dbUrl)

	// new migration instance
	m, err := migrate.New(sUrl, dbUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	// force migration to a specific version
	err = m.Force(migrationVersion)
	if err == migrate.ErrNoChange {
		log.Println("No change in database...")
		return nil
	} else if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Database migrations completed...")
	return nil
}
