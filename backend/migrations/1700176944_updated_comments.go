package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ccjrcenkq5jr4dk")
		if err != nil {
			return err
		}

		// update
		edit_attachments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gguxgjuf",
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
		}`), edit_attachments)
		collection.Schema.AddField(edit_attachments)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ccjrcenkq5jr4dk")
		if err != nil {
			return err
		}

		// update
		edit_attachments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
		}`), edit_attachments)
		collection.Schema.AddField(edit_attachments)

		return dao.SaveCollection(collection)
	})
}
