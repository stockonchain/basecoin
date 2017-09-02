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
	lasterr error;
}

func NewBasecoinDBPG() *BasecoinDBPG {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHostAddress, DBUser, DBPassword, DBName)
	con, err := sql.Open("postgres", dbinfo)
	checkErr(err);
	fmt.Println("Stockonchain connected to database")
	return &BasecoinDBPG{
		con : con,
		lasterr : err,
	}
}

func (db *BasecoinDBPG) Close() {
	db.con.Close();
}

func (db *BasecoinDBPG) AddTransaction() {
	var lastInsertId int
	db.lasterr = db.con.QueryRow("INSERT INTO TRANSACTION(AMOUNT, TYPE, PRODUCTID) VALUES($1,$2,$3) returning id;").Scan(&lastInsertId)
	checkErr(db.lasterr)
	fmt.Println("last inserted id =", lastInsertId)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}