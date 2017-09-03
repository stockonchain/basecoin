package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	DBHostAddress           = "localhost"
	DBPassword				= "postgres"
	DBUser					= "postgres"
	DBName					= "stockonchain"
)

type BasecoinDBPG struct {
	con *sql.DB;
}

func NewBasecoinDBPG() *BasecoinDBPG {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHostAddress, DBUser, DBPassword, DBName)
	con, err := sql.Open("postgres", dbinfo)
	checkErr(err);
	fmt.Println("Stockonchain connected to database")
	return &BasecoinDBPG{
		con : con,
	}
}

func (db *BasecoinDBPG) Close() {
	db.con.Close();
}

func (db *BasecoinDBPG) AddTransaction(amount int, rxTxType int, productId int) {
	var lastInsertId int
	err := db.con.QueryRow("INSERT INTO TRANSACTION(AMOUNT, TYPE, PRODUCTID) VALUES($1,$2,$3) returning TRANSACTIONID;", amount, rxTxType, productId).Scan(&lastInsertId)
	checkErr(err)
	if err != nil {
		fmt.Println("Transaction inserted with id =", lastInsertId)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}