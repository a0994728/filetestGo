package main

import (
	"fmt"
	"main/entity"
	create "main/file"
	"main/testDB"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("start")
	//DBに接続、正常に動かない場合はエラーが発生する。
	err := testDB.Db.Connection()
	if err != nil {
		fmt.Println(err, "connection error")
	}
	//DBの接続を切る
	defer testDB.Db.Close()
	//トランザクション処理を呼び出す。
	err = testDB.Db.Transaction(func(tx *sqlx.Tx) error {
		resultList := entity.WorldList{}
		err := tx.Select(&resultList, "SELECT * FROM tbl_country;")
		if err != nil {
			fmt.Println(err, "sql error")
		}
		//rowに格納された、SELECTを、forで1レコードずつ読み込み表示させる。

		// for rows.Next() {
		// 	world := entity.World{}
		// 	//rows.Scanの代わりにrows.StructScanを使う
		// 	err := rows.StructScan(&world)
		// 	if err != nil {
		// 		fmt.Println(err, "convert error")
		// 		// log.Fatal(err)

		// 	}
		// 	fmt.Println(world)
		// 	resultList = append(resultList, world)
		// }
		// fmt.Println("--------------------------")
		// for _, w := range resultList {

		// 	fmt.Println(w)
		// }
		err = create.CreateFile(resultList, "country.csv")
		return err
	})

	if err != nil {
		fmt.Println(err, "transaction failed")
	}
}
