package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
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
			"CREATE INDEX ` + "`" + `idx_xVbKRwQ` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `severity` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_8MwhI0j` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `author` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("by1itjt6pmxxvch")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_XSSe76p` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `category` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_oHJWfEM` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `severity` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xVbKRwQ` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `severity` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_8MwhI0j` + "`" + ` ON ` + "`" + `vulnerabilities` + "`" + ` (` + "`" + `author` + "`" + `)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
