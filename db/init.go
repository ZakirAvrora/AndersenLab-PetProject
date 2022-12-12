package db

import (
	"fmt"
	"github.com/ZakirAvrora/AndersenLab-PetProject/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"log"
)

func Init(conf *config.DatabaseConfig) *sqlx.DB {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)

	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln("database connection failed:", err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalln("database driver cannot be get:", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:../db/migrations", "postgres", driver)
	if err != nil {
		log.Fatalln("migration error:", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalln("migration error:", err)
	}

	return db
}
