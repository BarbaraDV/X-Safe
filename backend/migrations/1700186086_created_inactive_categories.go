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
			"id": "ux9enqs5s7l3xfq",
			"created": "2023-11-17 01:54:46.806Z",
			"updated": "2023-11-17 01:54:46.806Z",
			"name": "inactive_categories",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jmghfq1i",
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
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "select c.id, c.name, c.created, c.updated from categories as c\ninner join vulnerabilities as v\nwhere v.category != c.id\norder by c.name asc"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ux9enqs5s7l3xfq")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
