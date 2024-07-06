package database

import (
	"database/sql"

	"github.com/moha-abdi/telxchange/api/types/responses"
	"github.com/moha-abdi/telxchange/config"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	cfg := config.NewDBConfig()
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) AddPartner(partner responses.PartnerInfo) error {
	_, err := d.db.Exec(
		"INSERT INTO partner (id, name, subscriptionid, status) VALUES (?, ?, ?, ?)",
		partner.ID,
		partner.Name,
		partner.SubscriptionID,
		partner.Status,
	)
	return err
}
