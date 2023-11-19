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
		new_status := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kjit757p",
			"name": "status",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"pending",
					"reviewing",
					"solved"
				]
			}
		}`), new_status)
		collection.Schema.AddField(new_status)

		// update
		edit_severity := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
					"high",
					"critic"
				]
			}
		}`), edit_severity)
		collection.Schema.AddField(edit_severity)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kjit757p")

		// update
		edit_severity := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
		}`), edit_severity)
		collection.Schema.AddField(edit_severity)

		return dao.SaveCollection(collection)
	})
}
