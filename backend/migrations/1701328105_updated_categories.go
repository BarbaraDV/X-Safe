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

		collection, err := dao.FindCollectionByNameOrId("zjpl8d5al4dasan")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT DISTINCT v.category as id, v.category as name FROM vulnerabilities v"
		}`), &options)
		collection.SetOptions(options)

		// add
		new_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "npsrdobn",
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
		}`), new_name)
		collection.Schema.AddField(new_name)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("zjpl8d5al4dasan")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT DISTINCT v.category as id FROM vulnerabilities v"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("npsrdobn")

		return dao.SaveCollection(collection)
	})
}
