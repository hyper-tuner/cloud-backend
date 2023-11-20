package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "9eif9v40b0uw9l8",
			"created": "2023-11-20 10:50:08.919Z",
			"updated": "2023-11-20 10:50:08.919Z",
			"name": "iniFiles",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "h9yfwmvx",
					"name": "signature",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 3,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "z3e1a5cl",
					"name": "file",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"maxSize": 5242880,
						"mimeTypes": [
							"application/gzip",
							"application/octet-stream"
						],
						"thumbs": [],
						"protected": false
					}
				},
				{
					"system": false,
					"id": "t5uayom3",
					"name": "ecosystem",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"speeduino",
							"rusefi",
							"fome"
						]
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `officialIniFiles_created_idx` + "`" + ` ON ` + "`" + `iniFiles` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_unique_h9yfwmvx` + "`" + ` ON ` + "`" + `iniFiles` + "`" + ` (` + "`" + `signature` + "`" + `)"
			],
			"listRule": "",
			"viewRule": "",
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

		collection, err := dao.FindCollectionByNameOrId("9eif9v40b0uw9l8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
