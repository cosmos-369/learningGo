package poker

import (
	"testing"
)

func TestConnection(t *testing.T) {
	t.Run("check if we can connect to player store DB", func(t *testing.T) {
		db, err := connect()
		if db != nil {
			defer db.Close()
		}

		if err != nil {
			t.Errorf("unable to connect to DB got error:%q", err.Error())
		}
	})
}

func TestDBSchema(t *testing.T) {
	db, err := connect()
	if err != nil {
		db.Close()
		t.Errorf("unable to connect to DT got error:%q", err.Error())
	}
	defer db.Close()
	t.Run("check the schema of the DB", func(t *testing.T) {

		reqSchema := map[string]string{
			"id":   "bigint",
			"name": "text",
			"wins": "bigint",
		}

		rows, _ := db.Query("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = 'players'")
		defer rows.Close()
		for rows.Next() {
			var column_name string
			var data_type string
			if err = rows.Scan(&column_name, &data_type); err != nil {
				t.Error(err.Error())
			}

			if val, ok := reqSchema[column_name]; !ok || val != data_type {
				t.Errorf("the table does not contain the requried schema, got col:%q or type:%q, wanted schema %v",
					column_name, data_type, reqSchema)
			}
		}
	})
}
