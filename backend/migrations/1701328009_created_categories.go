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
			"id": "zjpl8d5al4dasan",
			"created": "2023-11-30 07:06:49.237Z",
			"updated": "2023-11-30 07:06:49.237Z",
			"name": "categories",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "gaeagbc4",
					"name": "category",
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
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT DISTINCT (ROW_NUMBER() OVER()) as id, v.category FROM vulnerabilities v"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("zjpl8d5al4dasan")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
