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

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		// add
		new_comments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "lgnqbijw",
			"name": "comments",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "ccjrcenkq5jr4dk",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_comments)
		collection.Schema.AddField(new_comments)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("lgnqbijw")

		return dao.SaveCollection(collection)
	})
}
