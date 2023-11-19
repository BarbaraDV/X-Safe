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
			"id": "by1itjt6pmxxvch",
			"created": "2023-10-31 16:13:15.291Z",
			"updated": "2023-10-31 16:13:15.291Z",
			"name": "vulnerabilities",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "akkywpob",
					"name": "title",
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
					"id": "ybosee8w",
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
					"id": "zshopt8l",
					"name": "attachments",
					"type": "file",
					"required": false,
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
					"id": "v42psnbj",
					"name": "author",
					"type": "relation",
					"required": false,
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
					"id": "dri78gkr",
					"name": "category",
					"type": "relation",
					"required": false,
					"presentable": true,
					"unique": false,
					"options": {
						"collectionId": "mytj1rirs8ldytn",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "g4bicprr",
					"name": "severity",
					"type": "select",
					"required": true,
					"presentable": false,
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
					"id": "wr96lqv8",
					"name": "incidents",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "r9h5q0s6i73watl",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_XSSe76p` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `category` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_oHJWfEM` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `severity` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_xVbKRwQ` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `severity` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_8MwhI0j` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `author` + "`" + `)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.id != ''",
			"updateRule": "@request.auth.id != ''",
			"deleteRule": "@request.auth.id != ''",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
