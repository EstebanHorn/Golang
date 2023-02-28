package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database
const url = "root:0704@tcp(localhost:3306)/blog_db"

var db *sql.DB

func ConnectDB() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("coneccion exitosa")
	db = connection
}

func DiscconectDb() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistTable(table string) bool {
	sql := fmt.Sprintf("SHOW TABLE LIKE '%s'", table)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("ERROR! ", err)
	}
	return rows.Next()
}

func CreateTable(shema string, tableName string) {
	if !ExistTable(tableName) {
		_, err := Exec(shema)
		if err != nil {
			fmt.Println("ERROR! ", err)
		}
	}

}

func ResetTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

func Exec(query string, args ...any) (sql.Result, error) {
	ConnectDB()
	res, err := db.Exec(query, args...)
	DiscconectDb()
	if err != nil {
		fmt.Println("ERROR! ", err)
	}
	return res, err
}

func Query(query string, args ...any) (*sql.Rows, error) {
	ConnectDB()
	rows, err := db.Query(query, args...)
	DiscconectDb()
	if err != nil {
		fmt.Println("ERROR! ", err)
	}
	return rows, err
}
