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
			"id": "mytj1rirs8ldytn",
			"created": "2023-10-31 16:13:15.291Z",
			"updated": "2023-10-31 16:13:15.291Z",
			"name": "categories",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "pwkuqdbm",
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
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.id != ''",
			"updateRule": "@request.auth.admin = true",
			"deleteRule": "@request.auth.admin = true",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("mytj1rirs8ldytn")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
