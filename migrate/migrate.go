package migrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

// Migrate func
func Migrate(migrationsSource string, db *sql.DB) error {
	data, err := ioutil.ReadFile(migrationsSource)
	if err != nil {
		return fmt.Errorf("error in migrations: %w", err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("error in migrations: %w", err)
	}

	return nil
}
