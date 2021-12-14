package migrations

import (
	"github.com/jmoiron/sqlx"
)

type Version20211214202145 struct {
}

func NewVersion20211214202145() Version20211214202145 {
	return Version20211214202145{}
}

func (v Version20211214202145) Up(tx *sqlx.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE "user" (
		    id       	UUID      NOT NULL,
		    roles      	JSONB,
		    username   	VARCHAR(255) NOT NULL,
		    password   	VARCHAR(255),
		    updated_at 	TIMESTAMP,
		    deleted_at  TIMESTAMP,
		    created_at 	TIMESTAMP NOT NULL DEFAULT (NOW()),
		    PRIMARY KEY (id),
		    UNIQUE (username)
		);
`)
	return err
}

func (v Version20211214202145) Down(tx *sqlx.Tx) error {
	_, err := tx.Exec(`DROP TABLE "user";`)
	return err
}
