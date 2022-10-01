package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,name) VALUES ('budi','Budi')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	// get koneksi database
	db := GetConnection()
	// tutup koneksi database di akhir perintah program
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	// tutup rows
	defer rows.Close()

	// ambil data dari rows
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySqlComplex(t *testing.T) {
	// get koneksi database
	db := GetConnection()
	// tutup koneksi database di akhir perintah program
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	// tutup rows
	defer rows.Close()

	// ambil data dari rows
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var created_at time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		fmt.Println("Email :", email)
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		fmt.Println("birthDate :", birthDate)
		fmt.Println("Married :", married)
		fmt.Println("Created_at :", created_at)
		fmt.Printf("============\n")
	}

	fmt.Println("Success insert new customer")
}
