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
			"CREATE INDEX ` + "`" + `idx_8MwhI0j` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `author` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_itgioL6` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `category` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_LHrV19W` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `incident` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_Bt3nqSX` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `trash` + "`" + `)"
		]`), &collection.Indexes)

		// add
		new_trash := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "knar9v1f",
			"name": "trash",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_trash)
		collection.Schema.AddField(new_trash)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
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
		collection.Schema.RemoveField("knar9v1f")

		return dao.SaveCollection(collection)
	})
}
