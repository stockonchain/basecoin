package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	//"errors"
	//sm "github.com/tendermint/basecoin/state"
	"github.com/tendermint/basecoin/types"
)

const (
	DBHostAddress           = "localhost"
	DBPassword				= "postgres"
	DBUser					= "postgres"
	DBName					= "stockonchain"
	DBTypeSendTxValue		= 0
	DBTypeSendRxValue		= 1
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

func AddTransactionSingle(db *BasecoinDBPG, rxTxType int) (int64, error) {
	var lastInsertId int64
	err := db.con.QueryRow("INSERT INTO transactionstock(type) VALUES($1) returning transactionid;", rxTxType).Scan(&lastInsertId)
	checkErr(err)
	if err != nil {
		return 0, err;
	}
	fmt.Println("Transaction inserted with id =", lastInsertId)
	return lastInsertId, err
}

func AddTransactionItemSingle(db *BasecoinDBPG, amount int64, trid int64, itemid int64) error {
	_, err := db.con.Query("INSERT INTO transaction_item(amount, transactionid, itemid) VALUES($1,$2,$3);", amount, trid, itemid)
	checkErr(err)
	if err != nil {
		return err;
	}
	return err
}


func (db *BasecoinDBPG) AddTransaction(ins []types.TxInput) error {
	for _, in := range ins {
		txn, err := db.con.Begin()
		if err != nil {
			return err;
		}
		var rxTxType int
		if in.Type {
			rxTxType = 1
		} else {
			rxTxType = 0
		}
		trid, err := AddTransactionSingle( db, rxTxType )
		if err != nil {
			return err;
		}
		for _, item := range in.Items {
			err := AddTransactionItemSingle(db, item.Amount, trid, item.ID )
			if err != nil {
				return err;
			}
		}
		txn.Commit()
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}