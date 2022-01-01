package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データストアから結果を取得するにはいくつか決まった方法があります。
    // ●行セットを返却するクエリを実行する。
    // ●繰り返し使用するステートメントを準備し、複数回実行し、破棄する。
    // ●繰り返し使用しない、1回限りのステートメントを実行する。
    // ●単一の行を返却するクエリを実行する。の方法がある
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}