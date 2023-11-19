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

		// update
		edit_severity := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
					"high",
					"critic"
				]
			}
		}`), edit_severity)
		collection.Schema.AddField(edit_severity)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("r9h5q0s6i73watl")
		if err != nil {
			return err
		}

		// update
		edit_severity := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
		}`), edit_severity)
		collection.Schema.AddField(edit_severity)

		return dao.SaveCollection(collection)
	})
}
