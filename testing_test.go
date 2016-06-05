package dbtt

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/driver/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tochti/speci"
)

func Test_IsInTable(t *testing.T) {
	specs, err := speci.ReadSQLite("test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(specs.Path)

	pool, err := specs.DB()
	if err != nil {
		t.Fatal(err)
	}

	db := sqlx.NewDb(pool, "sqlite3")
	defer db.Close()

	ResetDB(t, "sqlite3://"+specs.Path, "./test-fixtures")

	IsNotInTable(t, db, "test_table", "id=?", 1)

	_, err = db.Exec("INSERT INTO test_TABLE VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}

	IsInTable(t, db, "test_table", "id=?", 1)
}
