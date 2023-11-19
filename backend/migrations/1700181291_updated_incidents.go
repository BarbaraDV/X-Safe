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

		collection, err := dao.FindCollectionByNameOrId("r9h5q0s6i73watl")
		if err != nil {
			return err
		}

		// add
		new_comments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "p2x4iyd6",
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

		collection, err := dao.FindCollectionByNameOrId("r9h5q0s6i73watl")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("p2x4iyd6")

		return dao.SaveCollection(collection)
	})
}
