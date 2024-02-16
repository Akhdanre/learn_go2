package learngodatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
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

func TestQuerySelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id, "name :", name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_data, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birhtDate sql.NullTime
		var createAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birhtDate, &married, &createAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		fmt.Println("name :", name)
		if email.Valid {
			fmt.Println("email :", email)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birhtDate.Valid {
			fmt.Println("birthdate :", birhtDate)
		}
		fmt.Println("married :", married)
		fmt.Println("create at :", createAt)
		fmt.Println("=====================")
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "akhdan'; #"
	password := "akhdan"

	query := "SELECT username FROM user where username ='" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("berhasil Login", username)
	} else {
		fmt.Println("gagal Login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "akhdan"
	password := "salah"

	query := "SELECT username FROM user where username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("berhasil Login", username)
	} else {
		fmt.Println("gagal Login")
	}
}

func TestExecSQLParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "oukenze; DROP TABLE user; #"
	password := "lorem"

	query := "INSERT INTO user (username, password) values (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("succes insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "danakhdan12@gmail.com"
	comment := "lorem ipsum"

	query := "INSERT INTO comment (email, comment) values (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new comment with id", insertId)
}

func TestPrepareStmt(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comment (email, comment) values (?, ?)"
	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "dani" + strconv.Itoa(i) + "@gmail.com"
		comment := "data comment ke " + strconv.Itoa(i)

		res, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := res.LastInsertId()
		fmt.Println(lastInsertId)
	}
}
