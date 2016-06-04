package dbtt

import (
	"fmt"
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

	fmt.Println(specs.String())
	db, err := sqlx.Connect("sqlite3", specs.String())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ResetDB(t, specs.String(), "./test-fixtures")

	IsNotInTable(t, db, "test_table", "id=?", 1)

	_, err = db.Exec("INSERT INTO test_TABLE VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}

	IsInTable(t, db, "test_table", "id=?", 1)
}
