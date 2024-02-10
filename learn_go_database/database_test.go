package learngodatabase

import (
	"database/sql"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "dev:superdeveloper@tcp(localhost:3306)/learn_go_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
