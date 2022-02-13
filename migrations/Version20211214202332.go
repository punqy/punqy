package migrations

import (
	"github.com/jmoiron/sqlx"
)

type Version20211214202332 struct {
}

func NewVersion20211214202332() Version20211214202332 {
	return Version20211214202332{}
}

func (v Version20211214202332) Up(tx *sqlx.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE "oauth_client" (
		    id                  UUID      NOT NULL,
		    client_secret       VARCHAR(255),
		    allowed_grant_types JSONB              DEFAULT NULL,
		    updated_at          TIMESTAMP,
		    deleted_at          TIMESTAMP,
		    created_at          TIMESTAMP NOT NULL DEFAULT (NOW()),
		    PRIMARY KEY (id),
		    UNIQUE (client_secret)
		);
`)
	return err
}

func (v Version20211214202332) Down(tx *sqlx.Tx) error {
	_, err := tx.Exec(`DROP TABLE "oauth_client", "oauth_access_token", "oauth_refresh_token";`)
	return err
}
