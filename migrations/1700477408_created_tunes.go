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
			fmt.Println("collection tunes already exists")
			return nil
		}

		jsonData := `{
			"id": "5djmpehuiigg06b",
			"created": "2023-11-20 10:50:08.921Z",
			"updated": "2023-11-20 10:50:08.921Z",
			"name": "tunes",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "io2qgnvc",
					"name": "author",
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
					"id": "pkq4wfcj",
					"name": "tuneId",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 5,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "6gd6tzwx",
					"name": "source",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"web",
							"app"
						]
					}
				},
				{
					"system": false,
					"id": "jcjunqhl",
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
					"id": "lwbwtgmx",
					"name": "stars",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "g9b17t9y",
					"name": "vehicleName",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 2,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "w7qssd4t",
					"name": "engineMake",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 2,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "h47ir1bi",
					"name": "engineCode",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 2,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "frqzy24e",
					"name": "displacement",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "j1asw8n1",
					"name": "cylindersCount",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 16,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "0x8poyze",
					"name": "aspiration",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"na",
							"turbocharged",
							"supercharged"
						]
					}
				},
				{
					"system": false,
					"id": "ssd9iccu",
					"name": "compression",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "whk0u6fg",
					"name": "fuel",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "4ydhmn21",
					"name": "ignition",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 255,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "lti2l0im",
					"name": "injectorsSize",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100000,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "pmt4jrhm",
					"name": "year",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 2222,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "s9fjthhs",
					"name": "hp",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100000,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "s4csjkpt",
					"name": "stockHp",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100000,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "hfpkctpl",
					"name": "readme",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 5,
						"max": 3000,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "dk22rxkw",
					"name": "tags",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"base map",
							"help needed"
						]
					}
				},
				{
					"system": false,
					"id": "d5vpizsr",
					"name": "textSearch",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": 2048,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "1bjweixt",
					"name": "visibility",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"public",
							"unlisted"
						]
					}
				},
				{
					"system": false,
					"id": "roqbws6u",
					"name": "tuneFile",
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
					"id": "59syvgnj",
					"name": "customIniFile",
					"type": "file",
					"required": false,
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
					"id": "gzccot0l",
					"name": "logFiles",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 5,
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
					"id": "2z0i9ttc",
					"name": "toothLogFiles",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 5,
						"maxSize": 5242880,
						"mimeTypes": [
							"application/gzip",
							"application/octet-stream"
						],
						"thumbs": [],
						"protected": false
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `tunes_created_idx` + "`" + ` ON ` + "`" + `tunes` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_unique_pkq4wfcj` + "`" + ` ON ` + "`" + `tunes` + "`" + ` (` + "`" + `tuneId` + "`" + `)"
			],
			"listRule": "visibility = \"public\" || (visibility = \"unlisted\" && author = @request.auth.id)",
			"viewRule": null,
			"createRule": "@request.auth.verified = true",
			"updateRule": "@request.auth.id = author",
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

		collection, err := dao.FindCollectionByNameOrId("5djmpehuiigg06b")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
