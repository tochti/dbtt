package dbtt

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
)

type (
	FatalMethods interface {
		Fatal(args ...interface{})
	}
)

// Prüft ob Daten in Table gefunden werden
func isInTable(db *sqlx.DB, table, where string, args ...interface{}) (bool, error) {
	q := fmt.Sprintf(`SELECT COUNT(*) FROM %v WHERE %v`, table, where)
	q = db.Rebind(q)

	var count int
	err := db.QueryRow(q, args...).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func IsInTable(t FatalMethods, db *sqlx.DB, table, where string, args ...interface{}) {
	ok, err := isInTable(db, table, where, args...)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		err := fmt.Errorf("Is not in %v but it should", table)
		t.Fatal(err)
	}
}

func IsNotInTable(t FatalMethods, db *sqlx.DB, table, where string, args ...interface{}) {
	ok, err := isInTable(db, table, where, args...)
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		err := fmt.Errorf("Is in table %v but it shouldn't", table)
		t.Fatal(err)
	}
}

func ResetDB(t FatalMethods, url, migrationDir string) {
	errs, ok := migrate.ResetSync(url, migrationDir)
	if !ok {
		t.Fatal(errs)
	}
}
