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

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_oHJWfEM` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `severity` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_8MwhI0j` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `author` + "`" + `)"
		]`), &collection.Indexes)

		// remove
		collection.Schema.RemoveField("dri78gkr")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_oHJWfEM` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `severity` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xVbKRwQ` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `severity` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_8MwhI0j` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `author` + "`" + `)"
		]`), &collection.Indexes)

		// add
		del_category := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "dri78gkr",
			"name": "category",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "mytj1rirs8ldytn",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), del_category)
		collection.Schema.AddField(del_category)

		return dao.SaveCollection(collection)
	})
}
