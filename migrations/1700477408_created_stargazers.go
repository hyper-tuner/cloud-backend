package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
    		dao := daos.New(db);
		_, err := dao.FindCollectionByNameOrId("9eif9v40b0uw9l8")

		if err == nil {
			fmt.Println("collection stargazers already exists")
			return nil
		}

		jsonData := `{
			"id": "z8cojwcvlyxxyll",
			"created": "2023-11-20 10:50:08.921Z",
			"updated": "2023-11-20 10:50:08.921Z",
			"name": "stargazers",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "him7pbq2",
					"name": "user",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "ny7akrmn",
					"name": "tune",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "5djmpehuiigg06b",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `_z8cojwcvlyxxyll_created_idx` + "`" + ` ON ` + "`" + `stargazers` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `unique_stargazers_on_user_tune` + "`" + ` ON ` + "`" + `stargazers` + "`" + ` (` + "`" + `user` + "`" + `, ` + "`" + `tune` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("z8cojwcvlyxxyll")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
