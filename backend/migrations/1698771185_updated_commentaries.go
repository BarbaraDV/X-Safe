package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ccjrcenkq5jr4dk")
		if err != nil {
			return err
		}

		collection.Name = "comments"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ccjrcenkq5jr4dk")
		if err != nil {
			return err
		}

		collection.Name = "commentaries"

		return dao.SaveCollection(collection)
	})
}
