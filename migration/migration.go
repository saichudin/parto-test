package migration

import (
	"saichudin/parto-test/config"
	"saichudin/parto-test/contact"
)

func Migrate() {
	config.Db.AutoMigrate(contact.Contact{})
}