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
			"id": "r9h5q0s6i73watl",
			"created": "2023-10-31 16:13:15.291Z",
			"updated": "2023-10-31 16:13:15.291Z",
			"name": "incidents",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "wgsmmjhp",
					"name": "short_description",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "lxn9vjp2",
					"name": "description",
					"type": "editor",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				},
				{
					"system": false,
					"id": "dtq3j5ba",
					"name": "date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "ufy9jflo",
					"name": "severity",
					"type": "select",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"low",
							"medium",
							"high"
						]
					}
				},
				{
					"system": false,
					"id": "mu2igoux",
					"name": "author",
					"type": "relation",
					"required": false,
					"presentable": true,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
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

		collection, err := dao.FindCollectionByNameOrId("r9h5q0s6i73watl")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
