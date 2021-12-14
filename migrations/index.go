package migrations

import "github.com/slmder/migrate"

func Versions() migrate.Collection {
	return migrate.Collection{
		NewVersion20211214202145(),
		NewVersion20211214202332(),
	}
}
