package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...
		_, err := db.CreateUniqueIndex("stargazers", "unique_stargazers_on_user_tune", "user", "tune").Execute()
		if err != nil {
			return err
		}

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...
		_, err := db.DropIndex("stargazers", "unique_stargazers_on_user_tune").Execute()
		if err != nil {
			return err
		}

		return nil
	})
}
