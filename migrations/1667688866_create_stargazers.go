package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...
		db.CreateUniqueIndex("stargazers", "unique_stargazers_on_user_tune", "user", "tune").Execute()

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...
		db.DropIndex("stargazers", "unique_stargazers_on_user_tune").Execute()

		return nil
	})
}
