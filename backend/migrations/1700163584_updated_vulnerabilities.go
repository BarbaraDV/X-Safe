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

		// update
		edit_incident := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wr96lqv8",
			"name": "incident",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "r9h5q0s6i73watl",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_incident)
		collection.Schema.AddField(edit_incident)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		// update
		edit_incident := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
		}`), edit_incident)
		collection.Schema.AddField(edit_incident)

		return dao.SaveCollection(collection)
	})
}
