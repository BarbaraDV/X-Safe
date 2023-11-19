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

		collection, err := dao.FindCollectionByNameOrId("mytj1rirs8ldytn")
		if err != nil {
			return err
		}

		// update
		edit_subcategory := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "z8tjrmyl",
			"name": "subcategory",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "mytj1rirs8ldytn",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_subcategory)
		collection.Schema.AddField(edit_subcategory)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("mytj1rirs8ldytn")
		if err != nil {
			return err
		}

		// update
		edit_subcategory := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "z8tjrmyl",
			"name": "parent",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "mytj1rirs8ldytn",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_subcategory)
		collection.Schema.AddField(edit_subcategory)

		return dao.SaveCollection(collection)
	})
}
