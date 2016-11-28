package cmd

import (
	"log"

	"github.com/RichardKnop/go-oauth2-server/migrations"
	"github.com/RichardKnop/go-oauth2-server/oauth"
)

// Migrate runs database migrations
func Migrate() error {
	_, db, err := initConfigDB(true, false)
	log.Print(db)
	log.Print("NFIOANFUWA")
	log.Print(err)
	if err != nil {
		return err
	}
	defer db.Close()

	// Bootstrap migrations
	if err := migrations.Bootstrap(db); err != nil {
		return err
	}

	// Run migrations for the oauth service
	if err := oauth.MigrateAll(db); err != nil {
		return err
	}

	return nil
}
