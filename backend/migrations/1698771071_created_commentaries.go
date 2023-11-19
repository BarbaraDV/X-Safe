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
			"id": "ccjrcenkq5jr4dk",
			"created": "2023-10-31 16:51:11.178Z",
			"updated": "2023-10-31 16:51:11.178Z",
			"name": "commentaries",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "dgqyxmj1",
					"name": "author",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "gguxgjuf",
					"name": "attachments",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 99,
						"maxSize": 5242880,
						"mimeTypes": [],
						"thumbs": [],
						"protected": true
					}
				},
				{
					"system": false,
					"id": "ncmrbkw4",
					"name": "body",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [],
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

		collection, err := dao.FindCollectionByNameOrId("ccjrcenkq5jr4dk")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
