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
			"query": "SELECT DISTINCT v.category as id FROM vulnerabilities v"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("gaeagbc4")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("zjpl8d5al4dasan")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT DISTINCT (ROW_NUMBER() OVER()) as id, v.category FROM vulnerabilities v"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_category := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gaeagbc4",
			"name": "category",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_category)
		collection.Schema.AddField(del_category)

		return dao.SaveCollection(collection)
	})
}
