package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/driver"
	"github.com/verryp/gue-eco-test/internal/product/server"
)

func Run() {
	rootCmd := &cobra.Command{}

	cmds := []*cobra.Command{
		{
			Use:   "serve-http-product",
			Short: "Serve Product",
			Long:  "Run Product Services",
			Run: func(cmd *cobra.Command, args []string) {
				server.Start()
			},
		},
		{
			Use:   "migrate-up",
			Short: "Migrate Up DB",
			Long:  "Run all of your outstanding migrations",
			Run: func(cmd *cobra.Command, args []string) {
				conf, err := common.NewConfig()
				if err != nil {
					log.Logger.Err(err).Msgf("Config error | %v", err)
					panic(err)
				}

				src := getMigrateSrc()

				doMigrate(conf.DB, src, migrate.Up)
			},
		},
		{
			Use:   "migrate-down",
			Short: "Migrate Down DB",
			Long:  "Rollback all the migrations",
			Run: func(cmd *cobra.Command, args []string) {
				conf, err := common.NewConfig()
				if err != nil {
					log.Logger.Err(err).Msgf("Config error | %v", err)
					panic(err)
				}

				src := getMigrateSrc()

				doMigrate(conf.DB, src, migrate.Down)
			},
		},
	}

	rootCmd.AddCommand(cmds...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getMigrateSrc() migrate.FileMigrationSource {
	src := migrate.FileMigrationSource{
		Dir: "migrations/sql/product",
	}

	return src
}

func doMigrate(conf common.DB, mSource migrate.FileMigrationSource, direction migrate.MigrationDirection) error {
	db, err := driver.NewMySQLDatabase(conf)
	if err != nil {
		log.Err(err).Msg("error on db connection")
		return err
	}

	defer db.Db.Close()

	total, err := migrate.Exec(db.Db, "mysql", mSource, direction)
	if err != nil {
		log.Err(err).Msg("error when do migration")
		return err
	}

	log.Info().Msgf("Migrate success, total migrated: %d", total)
	return nil
}

func init() {
	cobra.OnInitialize()
}

func main() {
	Run()
}
