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

		// add
		new_deleted := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gjs6mcuw",
			"name": "deleted",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_deleted)
		collection.Schema.AddField(new_deleted)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ccjrcenkq5jr4dk")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("gjs6mcuw")

		return dao.SaveCollection(collection)
	})
}
