package dbtm

import (
	"aap/speci"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Test_IsInTable(t *testing.T) {
	specs, err := speci.ReadPostgreSQL("test")
	if err != nil {
		t.Fatal(err)
	}

	pool, err := specs.DB()
	if err != nil {
		t.Fatal(err)
	}

	db := sqlx.NewDb(pool, "postgres")

	ResetDB(t, specs.String(), "./test-fixtures")

	IsNotInTable(t, db, "test_table", "id=?", 1)

	_, err = db.Exec("INSERT INTO test_TABLE VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}

	IsInTable(t, db, "test_table", "id=?", 1)

}
