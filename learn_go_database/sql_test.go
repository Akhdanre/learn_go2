package learngodatabase

import (
	"context"
	"fmt"
	"testing"
)


func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) values ('robbani', 'robbani')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("succes insert new customer")
}