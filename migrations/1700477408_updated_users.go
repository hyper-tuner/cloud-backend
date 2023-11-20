package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("id = @request.auth.id && @request.data.verifiedAuthor = null")

		collection.DeleteRule = nil

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": false,
			"exceptEmailDomains": null,
			"manageRule": null,
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"requireEmail": true
		}`), &options)
		collection.SetOptions(options)

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `_systemprofiles0_created_idx` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `created` + "`" + `)"
		]`), &collection.Indexes)

		// remove
		collection.Schema.RemoveField("users_name")

		// add
		new_verifiedAuthor := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fqmcoapu",
			"name": "verifiedAuthor",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_verifiedAuthor)
		collection.Schema.AddField(new_verifiedAuthor)

		// update
		edit_avatar := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "users_avatar",
			"name": "avatar",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"maxSize": 5242880,
				"mimeTypes": [
					"image/jpg",
					"image/jpeg",
					"image/png",
					"image/svg+xml",
					"image/gif"
				],
				"thumbs": null,
				"protected": false
			}
		}`), edit_avatar)
		collection.Schema.AddField(edit_avatar)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("id = @request.auth.id")

		collection.ViewRule = types.Pointer("id = @request.auth.id")

		collection.UpdateRule = types.Pointer("id = @request.auth.id")

		collection.DeleteRule = types.Pointer("id = @request.auth.id")

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": true,
			"exceptEmailDomains": null,
			"manageRule": null,
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"requireEmail": false
		}`), &options)
		collection.SetOptions(options)

		json.Unmarshal([]byte(`[]`), &collection.Indexes)

		// add
		del_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "users_name",
			"name": "name",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_name)
		collection.Schema.AddField(del_name)

		// remove
		collection.Schema.RemoveField("fqmcoapu")

		// update
		edit_avatar := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "users_avatar",
			"name": "avatar",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"maxSize": 5242880,
				"mimeTypes": [
					"image/jpeg",
					"image/png",
					"image/svg+xml",
					"image/gif",
					"image/webp"
				],
				"thumbs": null,
				"protected": false
			}
		}`), edit_avatar)
		collection.Schema.AddField(edit_avatar)

		return dao.SaveCollection(collection)
	})
}
