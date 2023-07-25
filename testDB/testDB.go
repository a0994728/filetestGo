package testDB

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DBの型を作成
type MyDB struct {
	db *sqlx.DB
}

// MyDB型の変数DBとする。
var Db MyDB

// DBを接続するメソッド
func (m *MyDB) Connection() error {
	var err error
	//Open("ドライバー名,ユーザー名:ユーザーパスワード@(ip:port番号)/データベース名")
	m.db, err = sqlx.Open("mysql", "root:root@(127.0.0.1:3306)/worlddb")
	if err != nil {
		return err
	}
	return nil
}

// DB接続を切るメソッド
func (m *MyDB) Close() {
	if m.db != nil {
		m.db.Close()
	}
}

// トランザクション処理を行うメソッド
func (m *MyDB) Transaction(txFunc func(*sqlx.Tx) error) error {
	var err error
	//トランザクションの開始
	tx, err := m.db.Beginx()
	if err != nil {
		return err
	} //パニックが発生し、ロールバックされ、処理が強制終了する場合
	//エラーが検出され、ロールバックする場合
	//正常に作動しコミットする場合の3つに分かれる。
	defer func() {
		if p := recover(); p != nil {
			log.Println("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("rollback")
			tx.Rollback()
		} else {
			log.Println("commit")
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	//トランザクション処理の判定結果が返る。
	return err
}
